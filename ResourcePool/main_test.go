package main

import (
	"testing"
)

func Benchmark_useBytes(b *testing.B){
	poolSize := (1 << 10) //1K
	size := (1 << 15)     // 32K

	bp = NewBytePool(poolSize, size) //初始化byte pool，否则会panic

	for i :=0;i<b.N;i++{
		useBytesWithPool()
	}
}
