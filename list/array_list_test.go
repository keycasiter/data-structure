package list

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	list := NewArrayList()

	for i := 0; i < 10; i++ {
		list.Insert(0, i)
		list.Print()
	}
}
