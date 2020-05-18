package main

type bytePool struct {
	c     chan []byte // 用来进行同步的channel，相当于一个池子
	width int         // 每个byte数组的长度
}

func NewBytePool(size, width int) (bp *bytePool) {
	return &bytePool{
		c:     make(chan []byte, size),
		width: width,
	}
}

func (bp *bytePool) Get() (b []byte) {
	select { // 每次先从池子总拿bytes，如果拿不到，就重新生成
	case b = <-bp.c:
		//do nothing

	default:
		b = make([]byte, 0, bp.width)

	}
	return
}

func (bp *bytePool) Put(b []byte) {
	select {
	case bp.c <- b: // 每次归还的时候，如果池子中还有空间，就直接放回到池子中
		//return back
	default: // 否则直接丢弃
		//discard bytes
	}
	return
}
