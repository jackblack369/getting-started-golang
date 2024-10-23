package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.
func New(k int) *Tree {
	var add func(t *Tree, v int) *Tree
	add = func(t *Tree, v int) *Tree {
		if t == nil {
			return &Tree{Value: v}
		}
		if rand.Intn(2) == 0 {
			t.Left = add(t.Left, v)
		} else {
			t.Right = add(t.Right, v)
		}
		return t
	}

	rand.Seed(time.Now().UnixNano())
	var t *Tree
	for i := 1; i <= 10; i++ {
		t = add(t, i*k)
	}
	return t
}

func (t *Tree) String() string {
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return s
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}

	_, ok := <-ch2
	return !ok
}

func TestBinarySame(t *testing.T) {

}
