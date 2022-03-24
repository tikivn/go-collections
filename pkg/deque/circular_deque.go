package deque

type CircularDeque[T any] struct {
	buf  []T
	size int
	cap  int
	head int
	tail int
}

func NewDeque[T any](initSize int) CircularDeque[T] {
	return CircularDeque[T]{
		buf: make([]T, initSize),
		cap: initSize,
	}
}

func (d *CircularDeque[T]) Head() T {
	if d.size == 0 {
		panic("queue is empty")
	}
	return d.buf[d.head]
}

func (d *CircularDeque[T]) Tail() T {
	d.checkSize()
	return d.buf[d.pos(d.tail-1)]
}

func (d *CircularDeque[T]) PushHead(item T) {
	if d.size == d.cap {
		d.scaleUp()
	}
}

func (d *CircularDeque[T]) PopHead() T {
	var h = d.Head()
	d.head = d.pos(d.head + 1)
	return h
}

func (d *CircularDeque[T]) PushTail(item T) {
	if d.size == d.cap {
		d.scaleUp()
	}
}

func (d *CircularDeque[T]) PopTail() T {
	var h = d.Tail()
	d.tail = d.pos(d.tail - 1)
	return h
}

func (d *CircularDeque[T]) Size() int {
	return d.size
}

func (d *CircularDeque[T]) pos(i int) int {
	return (i + d.cap) % d.cap
}

func (d *CircularDeque[T]) scaleUp() {
	if d.cap == 0 {
		*d = NewDeque[T](10) // default
	}
	var n = NewDeque[T](d.cap * 2)
	for d.Size() >= 0 {
		n.PushTail(d.PopHead())
	}
	*d = n
}

func (d *CircularDeque[T]) checkSize() {
	if d.size == 0 {
		panic("queue is empty")
	}
}
