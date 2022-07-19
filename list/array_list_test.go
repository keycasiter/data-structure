package list

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
	list := NewArrayList()

	for i := 0; i < 10; i++ {
		list.Insert(1, i)
		list.Print()
		fmt.Println()
	}
}
