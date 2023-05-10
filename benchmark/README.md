# Golang benchmark

This project is an example of a golang benchmark.
A golang benchmark must follow these rules:
- The benchmark is executed thanks to the `testing.B` package
- It must contains functions with the `Benchmark<Something>(b *testing.B)` naming convention
- Each one of these function must have a cycle inside
- The benchmark output returns the medium time of each cycle iteration
- The benchmark is launched from the terminal with: `go test -bench=.`

## Output

This benchmark for MacBook Pro 13 M2 with go version: `go version go1.19 darwin/arm64` returns:
```bash
~/tests/benchmark  $ go test -bench=.
goos: darwin
goarch: arm64
pkg: benchmark
BenchmarkWith-8      	 1221086	       965.0 ns/op
BenchmarkWithout-8   	 2232382	       541.0 ns/op
PASS
ok  	benchmark	4.021s
```