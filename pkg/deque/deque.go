package deque

type Deque[T any] interface {
	Head() T
	Tail() T
	PushHead(T)
	PushTail(T)
	PopHead() T
	PopTail() T
}
