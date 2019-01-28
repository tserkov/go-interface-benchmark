# Go Interface Benchmark
After being challenged by a colleague on the performance impacts of using Go's typed interfaces, I wrote this benchmark to determine the (contrived) impacts of their use.

### Results
As run on my test machine, with an AMD Athlon II X2 270 @ 3.4GHz, using Go 1.11.
```
tserkov@github:~$ go version
go version go1.11 linux/amd64
tserkov@github:~$ go test -v -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/tserkov/go-interface-benchmark
BenchmarkWithAlloc-2                    2000000000               0.89 ns/op            0 B/op          0 allocs/op
BenchmarkWithFunc-2                     300000000                4.71 ns/op            0 B/op          0 allocs/op
BenchmarkWithPtrFunc-2                  1000000000               2.38 ns/op            0 B/op          0 allocs/op
BenchmarkWithCastAlloc-2                2000000000               0.89 ns/op            0 B/op          0 allocs/op
BenchmarkWithPtrCastAlloc-2             2000000000               0.88 ns/op            0 B/op          0 allocs/op
BenchmarkWithCastFunc-2                 1000000000               2.06 ns/op            0 B/op          0 allocs/op
BenchmarkWithCastPtrFunc-2              1000000000               2.38 ns/op            0 B/op          0 allocs/op
BenchmarkWithPtrCastFunc-2              1000000000               2.14 ns/op            0 B/op          0 allocs/op
BenchmarkWithPtrCastPtrFunc-2           1000000000               2.88 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutPtrAlloc-2              2000000000               0.59 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutPtrFunc-2               1000000000               2.11 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutPtrPtrFunc-2            500000000                2.38 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutNonPtrAlloc-2           2000000000               0.59 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutNonPtrFunc-2            1000000000               2.06 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutNonPtrPtrFunc-2         1000000000               2.86 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/tserkov/go-interface-benchmark       32.347s
```

### Analysis
For allocations, we can see that allocating structs into an interfaced-typed variable is slower.

As for actual function calls, well...

| Name | Variable Type | Function Receiver | Time |
| - | - | - | - |
| BenchmarkWithoutNonPtrFunc | `struct` | Non-pointer | 2.06ns |
| BenchmarkWithCastFunc | `struct`, cast from `interface` | Non-pointer | 2.06ns |
| BenchmarkWithoutPtrFunc | `*struct` | Non-pointer | 2.11ns |
| BenchmarkWithPtrCastFunc | `*struct`, cast from `interface` | Non-pointer | 2.14ns |
| BenchmarkWithoutPtrPtrFunc | `*struct` | Pointer | 2.38ns |
| BenchmarkWithCastPtrFunc | `struct`, cast from `interface` | Pointer | 2.38ns |
| BenchmarkWithPtrFunc | `interface` | Pointer | 2.38ns |
| BenchmarkWithoutNonPtrPtrFunc | `struct` | Pointer | 2.86ns |
| BenchmarkWithPtrCastPtrFunc | `*struct`, cast from `interface` | Pointer | 2.88ns |
| BenchmarkWithFunc | `interface` | Non-pointer | 4.71ns |

The slowest calls were those to a non-pointer receiver function on an interface-typed variable. Which is one of the most common signatures (ie. `net/http`).

The fastest in this benchmark was calling a non-pointer receiver function on a struct-typed variable.  However, in the real world, this would likely not be true in most non-trival cases, since the receiver struct is copied for the function.  There's also likely so compiler optimizations taking place here that I'm not aware of.

In reality, calling a pointer receiver function on a pointer to struct would probably be the fastest.

One discrepancy I'm not sure about is how a struct-typed variable, cast from an interface, with a pointer receiver function is somehow faster than the exact same call without the intermediate casting stage.

For a bit more context in what the CPU is doing even outside of Go's bytecode, be sure to check out [latency numbers every programmer should know](https://gist.github.com/jboner/2841832).

### Conclusion
In practice, if your usage is in a [hot spot](https://en.wikipedia.org/wiki/Hot_spot_(computer_programming)), you'll want to call functions on struct-typed variables. The performance impact of the function having a pointer receiver to the struct will depend on what the function is doing.

However, if your usage is _not_ in a hot spot, keep in mind that a nanosecond (ns) is _1 billionth of a second_. So even with the worst-case scenario of calling a non-pointer receiver function on an interface-typed variable, the cost is only ~4.7 billionths of a second. Don't let overzealous or premature optimization prevent your package from becoming readable or testable.
