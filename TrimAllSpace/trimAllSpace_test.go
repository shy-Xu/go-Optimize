package trimallspace

import (
	"testing"
)

/*
goos: windows
goarch: amd64
pkg: go-optimize/TrimAllSpace
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
Benchmark_TrimAllSpac1-12        9638870               124.7 ns/op            64 B/op          4 allocs/op
Benchmark_TrimAllSpace2-12      52010188                22.76 ns/op            0 B/op          0 allocs/op
PASS
ok      go-optimize/TrimAllSpace        4.489s
*/

const strWithSqpce = " 3+2  /4+5 +1"

func Benchmark_TrimAllSpac1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimAllSpace1(strWithSqpce)
	}
}

func Benchmark_TrimAllSpace2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimAllSpace2(strWithSqpce)
	}
}
