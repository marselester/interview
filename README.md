# Interview

Interview notes and Go implementation of a few problems from Elements of Programming Interviews.

Table of content:

- [Site reliability engineering](notes/sre.md)
- [Linux](notes/linux.md)
- [Consistency](notes/consistency.md)
- [Postgres](notes/postgres.md)
- [Probability theory](notes/prob.md)
- [Algorithms](https://github.com/marselester/alg)
- [Primitive types](#primitive-types)
  - [Bitwise operation](notes/bits.md)
- [Arrays](#arrays)
- [Linked lists](#linked-lists)
- [Sorting](#sorting)
- [Searching](#searching)
- [Tests](#tests)
- [Benchmarks](#benchmarks)

## Primitive types

Bitwise operation problems:

- **[count bits](primitive/count_bits.go)**
  — count the number of bits that are set to 1 in a non-negative integer
- **[parity](primitive/parity.go)**
  — compute parity of a non-negative integer (1 if set bits are odd)
- **[swap bits](primitive/swap_bits.go)**
  — swap i-th and j-th bits of an integer

## Arrays

Array problems:

- **[even odd](array/even_odd.go)**
  — reorder an array of integers so that even integers appear first
- **[Dutch flag](array/dutch_flag.go)**
  — reorder an array of integers so that integers smaller than pivot appear first,
    then integers that equal to pivot, and finally integers larger than pivot

## Linked lists

Linked list problems:

- **[reverse](linkedlist/reverse.go)**

## Sorting

Sorting problems:

- **[h-index](sorting/h_index.go)**
  — calculate h-index metric that measures both productivity and citation impact of a researcher

## Searching

Searching problems:

- **[sum 2020](searching/sum2020.go)** (advent of code)
  — find two/three numbers that add up to 2020

## Tests

Each solution is covered with a few tests.

```sh
$ go test ./primitive
```

Go EPI Judge helps to make sure solutions pass all the test cases from EPIJudge repository (csv files).

```sh
$ git clone https://github.com/stefantds/go-epi-judge.git
$ git clone https://github.com/adnanaziz/EPIJudge.git
$ cat > go-epi-judge/config.yml <<CFG
testDataFolder: $(pwd)/EPIJudge/test_data
runParallelTests: true
CFG
$ cd go-epi-judge
$ # Solution should be placed to ./epi/count_bits/solution.go.
$ go test ./epi/count_bits
```

## Benchmarks

Benchmarks help to compare performance of naive and improved solutions.
For example, [Bit Twiddling Hacks](https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetNaive)
shows various techniques to count bits set.

```sh
$ go test -bench=CountBits ./primitive/
BenchmarkCountBitsNaive-12       	138934572	         8.47 ns/op
BenchmarkCountBitsKernighan-12    	250800961	         4.79 ns/op
```

It is recommended to run benchmarks multiple times and check how stable they are
using [Benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool.

```sh
$ go get golang.org/x/perf/cmd/benchstat
$ BENCH_COUNT=10
$ go test -bench=CountBits -count=$BENCH_COUNT ./primitive/ | tee bench.txt
goos: darwin
goarch: amd64
BenchmarkCountBitsNaive-12       	142174494	         8.35 ns/op
BenchmarkCountBitsNaive-12       	143033582	         8.41 ns/op
BenchmarkCountBitsNaive-12       	142639047	         8.33 ns/op
BenchmarkCountBitsNaive-12       	143451723	         8.39 ns/op
BenchmarkCountBitsNaive-12       	143017363	         8.37 ns/op
BenchmarkCountBitsNaive-12       	143151940	         8.35 ns/op
BenchmarkCountBitsNaive-12       	143833944	         8.41 ns/op
BenchmarkCountBitsNaive-12       	143466840	         8.32 ns/op
BenchmarkCountBitsNaive-12       	142059639	         8.36 ns/op
BenchmarkCountBitsNaive-12       	143438925	         8.44 ns/op
BenchmarkCountBitsKernighan-12    	252099921	         4.75 ns/op
BenchmarkCountBitsKernighan-12    	247041158	         4.74 ns/op
BenchmarkCountBitsKernighan-12    	250844809	         4.73 ns/op
BenchmarkCountBitsKernighan-12    	251821089	         4.75 ns/op
BenchmarkCountBitsKernighan-12    	251136150	         4.81 ns/op
BenchmarkCountBitsKernighan-12    	250629847	         4.77 ns/op
BenchmarkCountBitsKernighan-12    	251043451	         4.74 ns/op
BenchmarkCountBitsKernighan-12    	251974099	         4.77 ns/op
BenchmarkCountBitsKernighan-12    	251576733	         4.75 ns/op
BenchmarkCountBitsKernighan-12    	250239159	         4.74 ns/op
$ benchstat bench.txt
name                  time/op
CountBitsNaive-12     8.37ns ± 1%
CountBitsKernighan-12  4.75ns ± 1%
```

Naive solution's mean is 8.37 nanoseconds with ± 1% variation across the samples.
Let's compare naive to improved solution:

- isolate the samples of each implementation to files `bench_a`, `bench_b`, `bench_...`
- give the samples the same name `X`

```sh
$ grep Benchmark bench.txt | sed 's/Benchmark[A-z]*/BenchmarkX/g' | split -l $BENCH_COUNT -a 1 - bench_
$ benchstat bench_a bench_b
name  old time/op  new time/op  delta
X-12  8.37ns ± 1%  4.75ns ± 1%  -43.21%  (p=0.000 n=10+10)
```

Naive approach is 43.21% slower.

References:

- [High performance Go workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#benchmarking)
- [Benchstat tips](https://github.com/golang/go/issues/23471)
