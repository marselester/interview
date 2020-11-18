# Linux

Table of content:

- [Process](#process)
- [Filesystem](#filesystem)
- [File I/O](#file-io)
- [Networking](#networking)

References:

- UNIX and Linux System Administration Handbook by Evi Nemeth, Garth Snyder, Trent Hein, Ben Whaley, Dan Mackin
- The Linux Programming Interface by Michael Kerrisk

CPU can operate in user mode and kernel mode.
Hardware instructions allow to switch between modes.
Areas of virtual memory can be marked as user space or kernel space.
When running in user mode, CPU can only access a user memory space.

Certain operations can be performed only while CPU is operating in kernel mode
(accessing memory-management hardware, initiating device I/O operations).

A system call is a controlled entry point into the kernel.
A syscall changes CPU mode from user to kernel, so CPU can access protected kernel memory.
The wrapper func copies syscall args to specific registers and
executes trap/sysenter machine instruction which causes CPU to switch to kernel mode.

## Process

A process consists of an address space (memory pages) and data structures within the kernel
(status, priority, owner, opened files/ports, used resources, signal mask).
Memory layout (segments):

- text — program instructions
- data — program's static vars
- stack — a piece of memory that grows and shrinks as functions are called and return
  and that is used to allocate storage for local vars
- heap — an area from which a program can dinamically allocate extra memory

A process has at least one thread (execution context within a process).
Each thread has its own stack and CPU context but operates within process's address space.
Threads often wait (sleep) for the kernel to complete a background work, e.g.,
when a thread reads from a file, the kernel must request disk blocks and
deliver them into process's address space.

An existing process must clone itself to create a new process using `fork()` system call.
The clone is largely identical to parent: inherits copies of data, stack, heap segments
(read-only text segment is shared), env vars, stdin, stdout, stderr file descriptors,
resource limit settings.
The clone can then exchange the program it is running with `execve()`
(destroys existing segments replacing with new segments from the new program).
If the original parent dies, init/systemd (process 1) becomes the new parent.

From child's point of view `fork()` returns zero.
Parent receives child's PID.
Parent and child examine the return value to figure out their role.

When process completes, it calls `_exit()` with exit code to notify the kernel.
Before a dead process can disappear, the kernel requires the parent to acknowledge its death (parent calls `wait()`).
The parent receives child's exit code and summary of child's resources use.

`mmap()` syscall can create:

- a file mapping which maps a region of a file into process's virtual memory
  (pages are automatically loaded from the file as required).
- anonymous mapping (pages initialized to zero)

The memory in one process's mapping may be shared with mappings in other processes
when processes map the same region of a file or
because a child process inherits a mapping from its parent.
When a mapping is created as shared, modifications are visible to other processes
and are carried through to the underlying file.

Signals are process-level interrupt requests, e.g., KILL terminates a process at the kernel level.
Signals can be sent:

- among processes as a means of communication
- by terminal driver to interrupt (ctrl+c INT) or suspend processes (ctrl+z TSTP)
- by kernel, e.g., division by zero
- by kernel to notify about death of a child or availability of data on I/O channel

A signal handler is called if a process has a designated handler.
Othwerwise the kernel takes some default action (terminate, generate core dumps).

`strace -p PID` displays signals the process receives and every system call it makes
(arguments and result code from the kernel).
The `-f` flag follows forked processes, `-e trace=file` displays only file-related operations.

Sytem load averages (`uptime`) quantify the average number of processes
that have been runnable over the previous 1, 5, and 15 minute intervals.
Busyness caused by I/O (disk traffic) is also taken into account in Linux.
For CPU bound systems, load averages should be less than the total number of CPU cores.
Otherwise the system is overloaded.

Kernel exposes info about the system in /proc directory (pseudo-filesystem)
where files' contents are created as files are read.
Process-specific info is divided into dirs named by PID, e.g., /proc/1/ describes init.

| file    | contents
| ---     | ---
| cgroup  | control groups to which the process belongs
| cmd     | command or program the process is executing
| cmdline | complete command line of the process
| cwd     | symlink to current dir of the process
| environ | the process environment variables
| exe     | symlink to the file being executed
| fd      | dir contains links for each open file descriptor
| fdinfo  | dir contains further info for each open file descriptor
| maps    | memory mapping (libraries a program is linked to or depends on)
| ns      | dir with links to each namespace used by the process
| root    | symlink to the process's root dir (set with chroot)
| stat    | general process status info (ps)
| statm   | memory usage info

Process's memory consumption is usually shown in KB:

- VIRT is amount of vurtual memory allocated including shared resources (libs)
- RES is portion of VIRT currently mapped to specific memory pages
- DATA reports data and stack segments

If a file is deleted from filesystem but is still referenced by a running process,
then `df` reports the space but `du` doesn't.
This disparity persists until the file descriptor is closed or the file is truncated.

## Filesystem

Filesystem's inode table contains one inode for each file.
Inodes are identified numerically by their sequential location in the inode table
(`ls -i` shows file's inode number).

Inode includes file type, UID, GID, access mode, timestamps (last access, modification, attr change),
number of hard links, size in bytes, number of blocks allocated to the file
(measured in units of 512-byte blocks), pointers to data blocks.

ext2 doesn't store the data blocks of a file contiguously or even in sequential order.
In ext2 each inode contains 15 pointers.
The first 12 point to the location of the first 12 data blocks.
The next points to a block of pointers that give the locations of the subsequent data blocks.
For a block size of 4 KB, max file size is 4 TB.

`file` command shows the type of a file:

1. regular file _-_
2. directory _d_
3. character device file _c_
4. block device file _b_
5. local domain socket _s_
6. named pipe (FIFO) _p_
7. symbolic link _l_

A file's name is stored within its parent directory, not with the file itself.
The filesystem can be arbitrarily deep, but each component of a pathname must have a name max 255 char long.
Max length of a path passed as syscall argument is 4095 bytes (Linux).

All hard links to the file are equivalent (have the same inode number, `find -inum <int>`).
They can't cross filesystem boundaries.
The filesystem doesn't release data blocks until its last hard link was deleted.

Symbolic (soft) link points to a file by name (path is stored as contents)
and it doesn't have any permission info of its own.

Filesystem passes requests to access device files to appropriate device drivers.
In `ls -l /dev/tty10` output 4 is major device number (tells the kernel which driver the file refers to),
and 10 is minor number (typically tells the driver which physical unit to address).

Socket is a connection between processes.
Local domain sockets are referred to through filesystem object rather than a network port.
They are created with `socket` syscall and removed with `unlink` syscall or `rm` command.
Named pipe serves similar purposes (allows communication between two local processes).

Twelve mode bits, are stored along with four bits of file-type info (can't be changed after file creation).
When the kernel runs an execucable file (e.g., `passwd`) with setuid/setgid bits,
it changes the effective UID/GID of the resulting process to the UID/GID of the file
rather than UID/GID of the user that ran the command.
When set on dir, setgid bit causes newly created files within the dir to
take on group ownership of the dir rather than default user group.
When sticky bit is set on a dir (e.g., /tmp/), the filesystem won't allow to delete or rename it.

Linux defines more file attributes (`lsattr`), e.g., `chattr -ai` makes a file append-only and immutable.

## File I/O

Using `dup()` syscall two file descriptors can refer to the same open file description
and share a file offset value.
Therefore, if the file offset is changed via one file descriptor (`read()`, `write()`, `lseek()`),
this change is visible through the other file descriptor.
For example, `2>&1` informs the shell to redirect stderr to stdout
by making file descriptor 2 a duplicate of file descriptor 1.

`open()` syscall:

- `O_CREAT` and `O_EXCL` flags provide atomic existence check and file creation
  (process ensures it's the creator of a file)
- `O_APPEND` flag provides atomic append (processes might append to a global log file)

## Networking

| layer       | description
| ---         | ---
| application | HTTP, DNS
| transport   | TCP, UDP
| network     | IP, ICMP
| data-link   | ARP, device drivers
| physical    | copper, optical fiber, radio waves

As a packet travels down (TCP -> IP -> Ethernet -> physical wire),
each protocol's packet becomes the payload part of the next protocol.
Packet's max length (MTU) is imposed by the link layer.

IPv4 splits packets to conform to the MTU of a particular network link.
TCP automatically does path MTU discovery, but not UDP.

For example, `ping -D -s 1472 google.com` sets "don't fragment" bit and 1472 of data bytes to be sent.
Resulting packet is 1500 bytes (IP header is 20 bytes, ICMP header is 8 bytes).
WireGuard adds 60 bytes [overhead](https://lists.zx2c4.com/pipermail/wireguard/2017-December/002201.html)
for IPv4, and 80 bytes for IPv6.

IP doesn't guarantee that packets will arrive in the order they were transmitted,
that they won't be duplicated, or even that they will arrive at all.
Nor does IP provide error recovery (packets with header errors are silently discarded).

Fragmented datagram can be reassembled at the destination only if all fragments arrive.
The entire datagram is unusable if any fragment is lost or contains transmission errors.

A network interface has a MAC address that distinguishes it on the physical network,
and one or more IP addresses.
Senders use ARP to discover mapping from IP to MAC addresses (`ip neigh`).

Routing info is stored in a table in the kernel (`netstat -rn`).
The kernel picks the most specific route (with the longest mask).
A host can route packets only to gateways reachable through a directly connected network.

Like IP, UDP is connectionless, it supplements IP with port numbers and a data checksum to detect errors.

TCP provides a reliable, connection-oriented, bidirectional, byte-stream channel between two endpoints.
For each endpoint the kernel maintains state info, send buffer, receive buffer.

Each TCP segment includes the sequence number of the first byte in the segment
(allows to maintain order and deduplicate segments).
When a TCP segment arrives, a receiver sends ack (including sequence number) to the sender.

If a segment arrives with errors, then it is discarded, and no ack is sent.
The sender starts a timer when each segment is transmitted.
If ack is not received in time, the segment is retransmitted.

TCP provides flow control (receiver sends a buffer space left in ack)
to prevent a fast sender from overwhelming a slow receiver,
and congestion control (transmission rate is adjusted based on packet loss)
to prevent a fast sender from overwhelming the network.
