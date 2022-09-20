package list

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList_InsertHead(t *testing.T) {
	l := NewSingleLinkedList(nil)
	for i := 1; i <= 10; i++ {
		l.InsertHead(&node{
			data: i,
		})
	}
	l.Print()
}

func TestSingleLinkedList_InsertTail(t *testing.T) {
	l := NewSingleLinkedList(nil)
	for i := 1; i <= 10; i++ {
		l.InsertTail(&node{
			data: i,
		})
	}
	l.Print()
}

func TestSingleLinkedList_Get(t *testing.T) {
	l := NewSingleLinkedList(nil)
	for i := 1; i <= 10; i++ {
		l.InsertTail(&node{
			data: i,
		})
	}
	l.Print()
	fmt.Println(l.GetIndex(&node{
		data: 10,
	}))
}

func TestSingleLinkedList_Delete(t *testing.T) {
	l := NewSingleLinkedList(nil)
	for i := 1; i <= 10; i++ {
		l.InsertTail(&node{
			data: i,
		})
	}
	l.Print()
	fmt.Println(l.Delete(&node{
		data: 10,
	}))
	l.Print()
}

func TestSingleLinkedList_Reverse(t *testing.T) {
	l := NewSingleLinkedList(nil)
	for i := 1; i <= 3; i++ {
		l.InsertTail(&node{
			data: i,
		})
	}
	l.Print()
	l.Reverse()
	l.Print()
}

func TestAppend(t *testing.T) {
	l := NewSingleLinkedList(nil)
	for i := 1; i <= 10; i++ {
		l.InsertTail(&node{
			data: i,
		})
	}
	l.Print()
	l.Append(&node{data: 10}, &node{data: 22})
	l.Print()
}

func TestPrepend(t *testing.T) {
	l := NewSingleLinkedList(nil)
	for i := 1; i <= 10; i++ {
		l.InsertTail(&node{
			data: i,
		})
	}
	l.Print()
	l.Prepend(&node{data: 10}, &node{data: 22})
	l.Print()
}
