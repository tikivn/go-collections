package deque

import (
	"math/rand"
	"testing"
)

// This test make sure CircularDeque and SimpleDeque have the same result
func Test_testing(t *testing.T) {
	var d1 = NewDeque[int](10)
	var d2 = SimpleDeque[int]{}
	for i := 0; i < 1000; i++ {
		var op = rand.Int() % 3
		switch op {
		case 0: // push
			if rand.Int()%2 == 0 {
				d1.PushHead(rand.Int())
				d2.PushHead(rand.Int())
			} else {
				d1.PushTail(rand.Int())
				d2.PushTail(rand.Int())
			}
		case 1: // pop
			if d1.Size() > 0 {
				if rand.Int()%2 == 0 {
					if i1, i2 := d1.PopHead(), d2.PopHead(); i1 != i2 {
						t.FailNow()
					}
				} else {
					if i1, i2 := d1.PopTail(), d2.PopTail(); i1 != i2 {
						t.FailNow()
					}
				}
			}
		case 2: // get
			if d1.Size() > 0 {
				if rand.Int()%2 == 0 {
					if i1, i2 := d1.Head(), d2.Head(); i1 != i2 {
						t.FailNow()
					}
				} else {
					if i1, i2 := d1.Tail(), d2.Tail(); i1 != i2 {
						t.FailNow()
					}
				}
			}
		}
	}
}
