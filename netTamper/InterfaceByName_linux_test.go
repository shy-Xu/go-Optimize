package netTamper

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -bench=InterfaceBy -benchmem -run=none
// optimize: 20% memory 5% CPU from output
/*
goos: linux
goarch: amd64
pkg: go-optimize/netTamper
cpu: Intel(R) Xeon(R) CPU E5-2680 v4 @ 2.40GHz
Benchmark_TamperedInterfaceByName-4   	    6182	    188662 ns/op	   18084 B/op	      32 allocs/op
Benchmark_NetInterfaceByName-4        	    5757	    193950 ns/op	   22840 B/op	      54 allocs/op
Benchmark_NetInterfaceByIndex-4       	    6046	    187412 ns/op	   16196 B/op	      28 allocs/op
PASS
ok  	go-optimize/netTamper	3.487s
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

func Benchmark_NetInterfaceByIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		net.InterfaceByIndex(1)
	}
}
