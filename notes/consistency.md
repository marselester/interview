# Consistency

Table of content:

- [Distributed locking](#distributed-locking)
- [Distributed transactions](#distributed-transactions)
  - [Derived data](#derived-data)
  - [Distributed sagas](#distributed-sagas)
- [Single-object consistency models](#single-object-consistency-models)
  - [Linearizability](#linearizability)
  - [Causal](#causal)
- [CRDT](#crdt)
- [Lamport timestamp](#lamport-timestamp)
- [Transactional consistency models](#transactional-consistency-models)

References:

- [How to do distributed locking](https://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html) by Martin Kleppmann
- Designing Data-Intensive Applications by Martin Kleppmann
- Database Internals by Alex Petrov
- [Consistency models](https://jepsen.io/consistency),
  [Strong consistency models](https://aphyr.com/posts/313-strong-consistency-models) by Kyle Kingsbury
- [Highly Available Transactions: Virtues and Limitations](http://www.bailis.org/papers/hat-vldb2014.pdf)
  by Peter Bailis, Aaron Davidson, Alan Fekete†, Ali Ghodsi, Joseph M. Hellerstein, Ion Stoica
- Kafka: The Definitive Guide by Neha Narkhede, Gwen Shapira, and Todd Palino
- [Distributed Sagas](https://speakerdeck.com/caitiem20/distributed-sagas-a-protocol-for-coordinating-microservices)
  by Caitie McCaffrey
- [MVCC in PostgreSQL](https://postgrespro.com/blog/pgsql/5967856) by Egor Rogov

## Distributed locking

Two clients want to read-modify-write a file on AWS S3 concurrently.
A distributed lock should prevent them from losing updates.
Due to a process pause or a packet delay a lock might expire causing unsafe changes.

The solution is to include a fencing token with every write request to the storage service.
A fencing token is a number that increases every time a client acquires the lock.
The storage service rejects write requests that have tokens lower than it's already seen
([caveat](https://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html#comment-4953488211)).

Fencing token can be:

- ZooKeeper's zxid or the znode version number
- message sequence number (append lock requests into log, e.g., Kafka)

## Distributed transactions

**Two-phase commit** (MySQL, Postgres):

1. propose — coordinator proposes the value and collects votes.
  If a node fails, coordinator cannot proceed with commit.
2. commit/abort — nodes make the result of the first phase visible

Coordinator and every node logs each step locally to recover.
If coordinator never recovers, its replacement has to collect votes for a given tx again.

**Calvin** (FaunaDB) removes coordination overhead during the execution phase
(all replicas get the same inputs and produce equivalent outputs)
by relying on deterministic tx order provided by sequencers
(they establish a global tx input sequence using Paxos or single-leader replication).
Sequencer:

- collects txs and groups them into short time window batches
- forwards the batch to the scheduler after the batch is replicated
- scheduler executes parts of tx in parallel while preserving the serial order specified by sequencer

**Spanner** (CockroachDB, YugaByteDB) uses Paxos for consistent tx log replication,
2PC over Paxos groups for cross-shard txs, and TrueTime for deterministic tx ordering.

**Percolator** (TiDB is based on Percolator model) provides tx on top of Bigtable using its conditional mutation API.
Each tx consults the timestamp oracle (source of monotonically increasing timestamps)
for a tx start timestamp, and during commit.
Writes are buffered and committed using a client-driven 2pc.

**RAMP** uses MVCC and metadata of in-flight operations to fetch any missing state updates from other nodes
(readers that overlap with a writer can be detected and repaired by retrieving required info).
Writes are installed and made visible using 2pc:

1. prepare — writes are placed to target partitions
2. commit/abort — writes are made available atomically across all partitions

### Derived data

Log-based systems based on deterministic retry and idempotency achieve consistency
similar to distributed transactions but might lack linearizability due to async updates.

Causally related events can be routed to the same partition,
[logical timestamps](#lamport-timestamp) can provide total order without coordination.
Read event and system state can be logged and referenced by id in other events to capture causal dependency.
[CRDTs](#crdt) can be used to resolve conflicts when events are delivered in arbitrary order.

### Distributed sagas

Distributed sagas is a protocol for coordinating requests among multiple services.
For example, a travel saga has three tasks (book a hotel, book a flight, book a car) which are implemented as separate services.
A saga coordinator knows about all those services
and executes/cancels tasks if one of the tasks failed.

The services must obey the following rules:

- requests must be idempotent.
  For example, book a hotel two times (request_id=1), only one request succeeds.
- a service must support a compensating request that semantically undoes the effect of a request.
  It cannot be aborted, must be idempotent as well.
- it doesn't matter in what order a coordinator cancels requests (commutative).
  For example, "cancel car, then cancel hotel" or "cancel hotel, then cancel car".

## Single-object consistency models

**Read your writes** means Bob can see own changes being saved (no promises about other users).
Solutions:

- read Bob's own writes from leader, e.g., his profile
- wait for replication lag interval after update to start reading from followers
- Bob provides a [logical timestamp](#lamport-timestamp) of his recent write, so backend will serve from up-to-date replica

**Monotonic reads** means that if Bob makes several reads in sequence,
he will not read old data after seeing new (no promises about other users).
Reading from a random replica (different replication lag) makes time appear to go backwards.
Solution: Bob should read from the same replica.

**Monotonic writes** means that if Bob wrote w1, then w2, then all users will see changes in the same order.
Note, there is no promise about total order of writes of all users, e.g.,
in partitioned db writes are not globally ordered.
Solutions:

- causally related writes should go to the same partition
- keep track of causal dependencies

**Writes follow reads** means Bob reads value from w1 and writes w2, then w2 must be visible after w1
(Bob's writes must logically follow his last read).
Solution: servers should wait to reveal new writes (buffering)
until each write's respective dependencies are visible on all replicas.

### Linearizability

Linearizability (total order, real-time) is a recency guarantee on reads and writes of a register.
It makes replicated data appear as though there was only a single copy,
and all operations on it are atomic (locks, CAS).

- sequential (linearizable writes, but reads might be stale if a store async updated from log)
- causal

Consensus algorithms (Zab, Raft) are linearizable (ZooKeeper and etcd provide linearizable writes,
request linearizable reads by calling ZooKeeper's _sync_ or etcd's quorum read).
They resemble single-leader replication with measures to prevent split brain and stale replicas.
Linearizable CAS and total order broadcast are both equivalent to consensus.

Single-leader replication is potentially linearizable if
writes go to the leader, reads go to the leader or sync replicated follower.
Violations: split brain (old leader serves requests), async replication (lost writes).
Linearizable storage can be built using total order broadcast
(no guarantee when message is delivered, one recipient may lag behind the others),
e.g., write an intent, read the log, check if the first intent is your own.
Total order broadcast requires reliable and totally ordered message delivery,
e.g., [Kafka](https://go-talks.appspot.com/github.com/marselester/kafka-for-gophers/kafka.slide):

- must wait for the original leader to come back online to prevent loss of committed messages
- 3 or 5 replicas for durability
- writes require >= 2 in-sync replicas
- waiting for all acks

Leaderless replication is probably not linearizable due to
"last write wins" conflict resolution (Cassandra),
sloppy quorums (quorum includes nodes which aren't designated to store a value).
Linearizable Dynamo-style quorums (performance penalty):

- quorum writes/reads `w + r > n` (up-to-date reads) can tolerate one unavailable node
  given n=3 replicas, at least w=2 nodes must confirm writes, and r=2 nodes must confirm reads
- reader must perform sync read repair before returning results
- writer must read the latest state of a quorum before sending its writes

Multi-leader replication is generally not linearizable because of conflicting writes:
multiple nodes concurrently process writes and async replicate them.

### Causal

Causality defines a partial order: some operations are ordered with respect to each other ("happens-before" relationship),
but some are incomparable (concurrent operations).
Effects of the causally related operations are visible in the same order to all processes.

- writes follow reads
- read your writes
- monotonic writes
- monotonic reads

Causal consistency is the strongest possible model that doesn't slow down due to network delays,
and remains available during network failures.

Uniqueness constraint requires to know when total order is finalized (e.g., no other node will claim a username).

## CRDT

Conflict-free replicated data type allows operations to be applied in any order without changing the result.
For example, a grow-only counter is updated by 3 nodes with initial state vectors.

    node 1       node 2       node 3
    [0, 0, 0]    [0, 0, 0]    [0, 0, 0]

Each node holds a last known counter update from all nodes.
Each node is allowed to modify its own counter in the vector.
For example, 1st and 3rd nodes incremented their counters.

    node 1       node 2       node 3
    [1, 0, 0]    [0, 0, 0]    [0, 0, 1]

Merge function is used to combine the results by picking the max counter for each slot.

    node 1 received state vector from node 3
    merge([1, 0, 0], [0, 0, 1]) = [1, 0, 1]

Current vector state is a sum of slot values.

    sum([1, 0, 1]) = 2

A counter that supports increments and decrements consists of two vectors (one for increments, another for decrements).

Grow-only set: each node can append elements to its local set.
Merging two sets is commutative.

Using two sets allows additions and removals.
Only values contained in the addition set can be added into the removal set.
Current state is `addset - remset`.

## Lamport timestamp

Lamport timestamp is a pair of (counter, node ID) that provides a total ordering of operations
consistent with causality without coordination:

- each node keeps a counter of operations it processed
- every node and every client keeps track of the max counter it has seen so far,
  and includes that max on every request
- when a node receives a request or response with a max counter greater than its own,
  it increases its own counter to that max

## Transactional consistency models

**Serializable** isolation means that transactions appear to have occurred in some serial order one after another
(in reality they may have run concurrently). Examples:

- literally executing txs in a serial order (single-threaded txs in Redis using Lua)
- two-phase locking (writers block readers and vice-versa)
- serializable snapshot isolation (Postgres, FoundationDB).
  SSI isn't linearizable because it doesn't include writes that are more recent than the snapshot.

**Snapshot** isolation prevents nonrepeatable read and lost update (Postgres, but not MySQL).
Readers don't block writers and vice-versa (conflicting writes are aborted).

**Cursor stability** isolation prevents lost updates.
Atomic operations are usually implemented by taking an exclusing lock on object when it's read
so that no other transaction can read it until the update has been applied.

**Read committed** isolation prevents dirty reads and writes (two reads within the same tx may see different results).

**Read uncommitted** isolation prevents dirty writes.

Dirty reads occur when Bob reads Alice's writes before she committed them.

Dirty writes occur when Bob overwrites Alice's not yet committed writes.

Lost update occurs when Bob and Alice concurrently perform read-modify-write operations,
because Bob's write doesn't include Alice's modification.
Solutions:

- row lock (select for update)
- advisory lock
- atomic writes with constraints
- Postgres repeatable read isolation automatically detects when a lost update has occurred
- force all atomic operations to be executed on a single thread
- CAS if a db supports it

| balance update  | tps  | latency    | txs processed
| ---             | ---  | ---        | ---
| advisory lock   | 2972 | 26.911 ms  | 800
| row lock        | 2610 | 30.642 ms  | 800
| broken          | 2418 | 33.081 ms  | 800
| constraint      | 1532 | 52.194 ms  | 406
| snapshot        | 572  | 139.783 ms | 81

<details>

Concurrently perform read-modify-write operations in Postgres to reproduce lost update anomaly.

```sh
psql -U postgres -c '
create table account (
    id bigserial primary key,
    balance numeric not null default 0
);
insert into account select from generate_series(1, 10);'

# Broken balance update.
cat > lostupdate.sql <<SQL
\set r random_zipfian(1, 10, 1.07)
begin;
do \$\$ begin
    if (select balance from account where id=:r) >= 1 then
        update account set balance = balance - 1 where id=:r;
    end if;
end \$\$;
commit;
SQL
# "for update" implies that a row is fully updated (or deleted).
cat > lostupdate-rowlock.sql <<SQL
\set r random_zipfian(1, 10, 1.07)
begin;
do \$\$ begin
    if (select balance from account where id=:r for update) >= 1 then
        update account set balance = balance - 1 where id=:r;
    end if;
end \$\$;
commit;
SQL
# Transaction level advisory lock.
cat > lostupdate-advisorylock.sql <<SQL
\set r random_zipfian(1, 10, 1.07)
begin;
do \$\$ begin
    perform pg_advisory_xact_lock(:r);
    if (select balance from account where id=:r) >= 1 then
        update account set balance = balance - 1 where id=:r;
    end if;
end \$\$;
commit;
SQL
# Transaction with repeatable read isolation level.
cat > lostupdate-snapshot.sql <<SQL
\set r random_zipfian(1, 10, 1.07)
begin transaction isolation level repeatable read;
do \$\$ begin
    if (select balance from account where id=:r) >= 1 then
        update account set balance = balance - 1 where id=:r;
    end if;
end \$\$;
commit;
SQL

for f in lostupdate*; do
    psql -U postgres -c 'update account set balance=100;'
    pgbench -U postgres --client=80 -f $f 2> /dev/null
    psql -U postgres -c 'select * from account where balance!=0;'
done

# Overdraft constraint.
psql -U postgres -c 'alter table account add constraint overdraft check (balance >= 0);'
psql -U postgres -c 'update account set balance=100;'
pgbench -U postgres --client=80 -f lostupdate.sql 2> /dev/null
psql -U postgres -c 'select * from account where balance!=0;'
```

</details>

Nonrepeatable read occurs when two reads within the same tx get different results (reads violate causality).
Solutions: snapshot isolation.

Write skew occurs when two txs modify disjoint sets of values and violate invariants (two on-call doctors).
Solutions:

- lock rows
- multiple rows constraint or triggers
- serializable

Phantom reads occur when tx reads rows matching a search condition and
another write affects the result of that search (meeting room booking).
Solutions:

- materializing conflicts
- predicate or index-range locks
- serializable
