package list

import "fmt"

type ArrayList struct {
	elements [100]int
	size     int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		elements: [100]int{},
		size:     0,
	}
}

func (list *ArrayList) GetSize() int {
	return list.size
}

func (list *ArrayList) Insert(idx int, element int) bool {
	if idx < 1 || idx > list.size+1 {
		fmt.Println("Insert failed, index range error.")
		return false
	}
	if len(list.elements) == list.size {
		fmt.Println("Insert failed, space is full.")
		return false
	}

	for i := list.size - 1; i >= idx-1; i-- {
		list.elements[i+1] = list.elements[i]
	}
	list.elements[idx-1] = element
	list.size++

	return true
}

func (list *ArrayList) Print() {
	for i := 0; i < list.size; i++ {
		fmt.Printf("%d ", list.elements[i])
	}
}
