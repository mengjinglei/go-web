package main

var bp *bytePool

func useBytesWithPool() {
	data := bp.Get()
	_ = data //do some opration
	bp.Put(data)
}

func useBytes() {
	size := (1 << 15) // 32K
	data := make([]byte, 0, size)
	_ = data //do some opration
}

func main() {
	poolSize := (1 << 10) //1K
	size := (1 << 15)     // 32K

	bp = NewBytePool(poolSize, size)

	return
}
