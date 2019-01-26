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

# T Alloc
BenchmarkWithAlloc-2                    2000000000               0.89 ns/op            0 B/op          0 allocs/op
BenchmarkWithCastAlloc-2                2000000000               0.89 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutPtrAlloc-2              2000000000               0.89 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutNonPtrAlloc-2           2000000000               0.89 ns/op            0 B/op          0 allocs/op

# T.F()
BenchmarkWithFunc-2                     300000000                4.16 ns/op            0 B/op          0 allocs/op
BenchmarkWithCastFunc-2                 2000000000               0.63 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutPtrFunc-2               2000000000               0.61 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutNonPtrFunc-2            2000000000               0.59 ns/op            0 B/op          0 allocs/op

# T.PtrF()
BenchmarkWithPtrFunc-2                  1000000000               2.39 ns/op            0 B/op          0 allocs/op
BenchmarkWithCastPtrFunc-2              2000000000               0.60 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutPtrPtrFunc-2            2000000000               0.60 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutNonPtrPtrFunc-2         2000000000               0.60 ns/op            0 B/op          0 allocs/op

PASS
ok      github.com/tserkov/go-interface-benchmark       19.487s
```
The output has been reorganized to keep the comparative benchmarks together for readability.

### Analysis
For allocation of `T`, we can see no impact between the different methods, even when casting the `I`-typed variable into a second `T`-typed variable.

Immediately, we can see a 3.98x slowdown at 2.39ns/op when calling `T.PtrF()` (pointer receiver) and a 7.05x slowdown at 4.16ns/op when calling `T.F()` (non-pointer receiver) _from the `I`-typed variable_. This demonstrates a cost in calling functions through interface-typed variables _without_ first casting to a struct-typed variable.

There was some interesting minor variance (6%) between calling `T.F()`, ranging from 0.59ns/op to 0.63ns/op.

Calls to the struct-typed pointer-receiver function were very consistent at 0.60ns/op.

For a bit more context in what the CPU is doing even outside of Go's bytecode, be sure to check out [latency numbers every programmer should know](https://gist.github.com/jboner/2841832).

### Conclusion
In practice, if your usage is in a [hot spot](https://en.wikipedia.org/wiki/Hot_spot_(computer_programming)), you'll want to call functions on struct-typed variables. The performance impact of the function having a pointer receiver to the struct will depend on what the function is doing. For these no-op benchmarks, they performed equally.

However, if your usage is _not_ in a hot spot, keep in mind that a nanosecond (ns) is _1 billionth of a second_. So even with the worst-case scenario of calling a non-pointer-receiver function on an interface-typed variable, the cost is only ~4 billionths of a second. Don't let overzealous or premature optimization prevent your package from becoming readable or testable.
