package list

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestOrderList_Prepend(t *testing.T) {
	list := NewOrderList(10)
	for i := 0; i < 10; i++ {
		assert.Equal(t, true, list.Prepend(i+1))
	}
	for i := 0; i < 10; i++ {
		assert.Equal(t, false, list.Prepend(i+1))
	}
	list.Print()
}

func TestOrderList_Append(t *testing.T) {
	list := NewOrderList(10)
	for i := 0; i < 10; i++ {
		assert.Equal(t, true, list.Append(i+1))
	}
	for i := 0; i < 10; i++ {
		assert.Equal(t, false, list.Append(i+1))
	}
	list.Print()
}

func TestOrderList_Delete(t *testing.T) {
	list := NewOrderList(10)
	for i := 0; i < 10; i++ {
		assert.Equal(t, true, list.Append(i+1))
	}
	assert.Equal(t, true, list.Delete(1))
	for i := 0; i < 9; i++ {
		assert.Equal(t, true, list.Delete(0))
		list.Print()
	}
	assert.Equal(t, 0, list.Length)

}
