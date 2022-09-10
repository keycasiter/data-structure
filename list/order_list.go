package list

import "fmt"

type OrderList struct {
	//存储空间
	Elems []int
	//使用
	Length int
	//空间
	Size int
}

func NewOrderList(size int) *OrderList {
	return &OrderList{
		Elems: make([]int, 10),
		Size:  size}
}

func (l *OrderList) Prepend(x int) bool {
	if l.Length >= l.Size {
		return false
	}
	//所有元素向后移动
	target := l.Elems[0]
	for i := 0; i < l.Length; i++ {
		next := l.Elems[i+1]
		l.Elems[i+1] = target
		target = next
	}
	//前插新值
	l.Elems[0] = x
	//更新len计数
	l.Length++

	return true
}

func (l *OrderList) Append(x int) bool {
	if l.Length >= l.Size {
		return false
	}
	//前插新值
	l.Elems[l.Length] = x
	//更新len计数
	l.Length++
	return true
}

func (l *OrderList) Delete(index int) bool {
	if index < 0 || index > l.Length-1 || l.Length == 0 {
		return false
	}
	//处理最后一个
	if index == l.Length-1 {
		l.Elems[index] = 0
		l.Length--
		return true
	}
	//x位置之后所有元素向前移动
	for i := index; i < l.Length-1; i++ {
		next := l.Elems[i+1]
		l.Elems[i] = next
		l.Elems[i+1] = 0
	}
	//更新len计数
	l.Length--
	return true
}

func (l *OrderList) Print() {
	for _, each := range l.Elems {
		fmt.Printf("[%d]", each)
	}
	fmt.Println()
}
