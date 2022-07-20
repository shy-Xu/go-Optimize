package netTamper

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -bench=InterfaceBy -benchmem -run=none
// optimize: 28% memory 3% CPU from output
/*
goos: linux
goarch: amd64
pkg: go-optimize/netTamper
cpu: Intel(R) Xeon(R) CPU E5-2680 v4 @ 2.40GHz
Benchmark_TamperedInterfaceByName-4   	    6092	    187719 ns/op	   16068 B/op	      26 allocs/op
Benchmark_NetInterfaceByName-4        	    5724	    194584 ns/op	   22840 B/op	      54 allocs/op
Benchmark_NetInterfaceByIndex-4       	    6252	    187021 ns/op	   16196 B/op	      28 allocs/op
PASS
ok  	go-optimize/netTamper	3.498s
*/

func Test_TamperedInterfaceByName(t *testing.T) {
	res, err := TamperedInterfaceByName("eth0")
	assert.Nil(t, err)
	expect, err := net.InterfaceByName("eth0")
	assert.Nil(t, err)
	assert.Equal(t, expect, res)
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
