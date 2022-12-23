# Systems performance

## Amdahl's law

Amdahl's law shows the effectiveness of improving the performance
of one part of the system.

Let's say a task requires `T = 25 hours` to finish.
Its subtask takes 60% of time, `α = 0.6` which is `α * T = 15 hours`.

The subtask was optimized and now it takes x1.5 less time, `k = 1.5`
(the performance improvement factor).
The subtask now takes `(α * T) / k = 15 hours / 1.5 = 10 hours`.

The overall new time to finish a task is `T - α*T + (α*T)/k` or
`T * (1 - α + α/k)`.

What is the speed up for the whole task?
It's measured as `S = T_old / T_new`.

```
            T
S = ----------------- = 1 - α + α/k
    T * (1 - α + α/k)

            1
S = ----------------- = 1 / 0.8 = 1.25
    1 - 0.6 + 0.6/1.5
```

Even though the subtask runs x1.5 faster, the overall task only got x1.25 of speed up.

## Little's law

See Little's law in action at [capacity management](https://github.com/marselester/capacity).

`N = X * R`:

- N is capacity (number of workers)
- X is throughput (requests arrival rate)
- R is service time (how long it takes a worker to process a request)

The server has 7 workers, each takes 1 second on average to process a request.
The server should be able to handle 7 requests per second.

```
N = X * R
7 workers = X rps * 1s
X = 7/1 = 7 rps
```

## Universal scalability law

See [Web service scalability](https://go-talks.appspot.com/github.com/marselester/scalability/scalability.slide) slides.

![](https://raw.githubusercontent.com/marselester/scalability/master/img/usl.png)

- N is a number of nodes
- X(N) is throughput of the system of N nodes, e.g., 100 requests per second
- λ is a throughput of the system with one node X(1)
- σ is a coefficient of contention, e.g., σ=0.05 (5%)
- κ is a coefficient of crosstalk, e.g., κ=0.02 (2%)
