package deque

// SimpleDeque is an trivial implementation of Deque
type SimpleDeque[T any] struct {
	buf []T
}

func (t *SimpleDeque[T]) Head() T {
	return t.buf[0]
}

func (t *SimpleDeque[T]) Tail() T {
	return t.buf[len(t.buf)-1]
}

func (t *SimpleDeque[T]) PushHead(i T) {
	var temp = t.buf
	t.buf = append(make([]T, 0, len(t.buf)+1), i)
	for _, i := range temp {
		t.buf = append(t.buf, i)
	}
}

func (t *SimpleDeque[T]) PopHead() T {
	var item = t.Head()
	t.buf = t.buf[1:]
	return item
}

func (t *SimpleDeque[T]) PushTail(i T) {
	t.buf = append(t.buf, i)
}

func (t *SimpleDeque[T]) PopTail() T {
	var item = t.Tail()
	t.buf = t.buf[:len(t.buf)-1]
	return item
}

func (t *SimpleDeque[T]) Size() int {
	return len(t.buf)
}
