# crypto-benchmark

This test compares the performance of
- ECDSA VS Ed25519
- SHA256 VS Blake2b

To execute the benchmark run:

```bash
go test -bench=.
```

The following is the output

```
goos: darwin
goarch: amd64
pkg: crypto-benchmark
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkECDSAP256KeyGeneration-12         79281             14627 ns/op
BenchmarkECDSAP256Sign-12                  50302             23848 ns/op
BenchmarkECDSAP256Verify-12                16570             73127 ns/op
BenchmarkEd25519KeyGeneration-12           75586             15678 ns/op
BenchmarkEd25519Sign-12                    72140             16342 ns/op
BenchmarkEd25519Verify-12                  18495             64799 ns/op
BenchmarkSHA256-12                       5845134               206.3 ns/op
BenchmarkBlake2-12                       8091025               147.5 ns/op
PASS
ok      crypto-benchmark        12.330s
```