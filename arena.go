package dht

// arena is a free list that provides quick access to pre-allocated byte
// slices, greatly reducing memory churn and effectively disabling GC for these
// allocations. After the arena is created, a slice of bytes can be requested by
// calling Pop(). The caller is responsible for calling Push(), which puts the
// blocks back in the queue for later usage. The bytes given by Pop() are *not*
// zeroed, so the caller should only read positions that it knows to have been
// overwitten. That can be done by shortening the slice at the right place,
// based on the count of bytes returned by Write() and similar functions.
type arena chan []byte//这应该是模拟堆栈的功能

func newArena(blockSize int, numBlocks int) arena {
	blocks := make(arena, numBlocks)
	for i := 0; i < numBlocks; i++ {
		blocks <- make([]byte, blockSize)//4096字节 为4k 预先占用内存以避免后面申请内存和gc回收造成的资源浪费
	}
	return blocks
}

func (a arena) Pop() (x []byte) {//弹出
	return <-a
}

func (a arena) Push(x []byte) {//压入
	x = x[:cap(x)]
	a <- x
}
