## commonds benchmark golang

### run all benchmark
```bash
go test -bench=.
```

### run specifically name benchmark
```bash
go test -bench=BenchmarkFuncName
```

### run benchmark show profile CPU and memory
```bash
go test -bench=. -cpuprofile=cpu.pprof -memprofile=mem.pprof
```
 example output bechmark in golang
 ```bash
 BenchmarkSum-8    10000000        0.0750 ns/op
``

detail information
-BenchmarkSum-8: Nama benchmark dan jumlah core CPU yang digunakan.
-10000000: Jumlah iterasi yang dilakukan.
-0.0750 ns/op: Waktu rata-rata per iterasi dalam nanodetik. 
