package netTamper

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -bench=InterfaceByName -benchmem -run=none
// optimize: 20% memory 5% CPU from output
/*
goos: linux
goarch: amd64
pkg: test
cpu: Intel(R) Xeon(R) CPU E5-2680 v4 @ 2.40GHz
Benchmark_TamperedInterfaceByName-4   	    6160	    187941 ns/op	   18152 B/op	      34 allocs/op
Benchmark_NetInterfaceByName-4        	    5852	    194586 ns/op	   22840 B/op	      54 allocs/op
PASS
ok  	test	2.346s
*/

func Test_TamperedInterfaceByName(t *testing.T) {
	res, err := TamperedInterfaceByName("eth0")
	assert.Nil(t, err)
	expect, err := net.InterfaceByName("eth0")
	assert.Nil(t, err)
	assert.Equal(t, res, expect)
}

func Benchmark_TamperedInterfaceByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TamperedInterfaceByName("eth0")
	}
}

func Benchmark_NetInterfaceByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		net.InterfaceByName("eth0")
	}
}
