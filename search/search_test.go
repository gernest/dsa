package search

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	sample := []int{0, 1, 2, 8, 13, 17, 19, 32, 42}
	if Binary(sample, 3) {
		t.Error("expected false got true")
	}
	if !Binary(sample, 13) {
		t.Error("expected true got false")
	}
}

func TestBinaryTreePut(t *testing.T) {
	sample := []struct {
		key int
		val string
	}{
		{5, "five"},
		{30, "thirty"},
		{2, "two"},
		{40, "forty"},
		{24, "twentry five"},
		{4, "four"},
	}
	b := &BinaryTree{}
	for _, v := range sample {
		b.Put(v.key, v.val)
	}
	fmt.Println(b.Length())
	fmt.Println(b.root.Collect())
}
