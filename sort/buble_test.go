package sort

import (
	"reflect"
	"testing"
)

func TestBuble(t *testing.T) {
	src := []int{3, 11, 2, 9, 1, 5}
	expct := []int{1, 2, 3, 5, 9, 11}
	Buble(src)

	if !reflect.DeepEqual(src, expct) {
		t.Errorf("expecte %v to equal %v", src, expct)
	}
}
