package list

import "fmt"

type node struct {
	data int
	next *node
}

type SingleLinkedList struct {
	Head *node
}

func NewSingleLinkedList(head *node) *SingleLinkedList {
	return &SingleLinkedList{Head: head}
}

//头节点插入
func (l *SingleLinkedList) InsertHead(node *node) bool {
	if l.Head == nil {
		l.Head = node
		return true
	}

	cur := l.Head
	node.next = cur
	l.Head = node
	return true
}

//尾节点插入
func (l *SingleLinkedList) InsertTail(node *node) bool {
	if l.Head == nil {
		l.Head = node
		return true
	}
	tail := l.Head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = node
	return true
}

//指定节点之后插入
func (l *SingleLinkedList) Append(targetNode *node, newNode *node) bool {
	if l.Head == nil {
		return false
	}
	cur := l.Head
	for cur != nil {
		if cur.data == targetNode.data {
			next := cur.next
			cur.next = newNode
			newNode.next = next
			return true
		}

		cur = cur.next
	}
	return false
}

//指定节点之前插入
func (l *SingleLinkedList) Prepend(targetNode *node, newNode *node) bool {
	if l.Head == nil {
		return false
	}
	//头节点处理
	if l.Head.data == targetNode.data {
		newNode.next = l.Head
		l.Head = newNode
	}

	prev := l.Head
	cur := l.Head.next
	for cur != nil {
		if cur.data == targetNode.data {
			prev.next = newNode
			newNode.next = cur
			return true
		}

		prev = cur
		cur = cur.next
	}
	return false
}

// 获取节点位置
func (l *SingleLinkedList) GetIndex(node *node) int {
	idx := 0
	cur := l.Head
	for cur != nil {
		if cur.data == node.data {
			return idx
		}
		cur = cur.next
		idx++
	}

	return -1
}

// 删除指定节点
func (l *SingleLinkedList) Delete(node *node) bool {
	if l.Head == nil {
		return false
	}
	prev := l.Head
	cur := l.Head
	for prev != nil && cur != nil {
		if cur.data == node.data {
			//头节点处理
			if cur == l.Head {
				l.Head = cur.next
				return true
			}

			prev.next = cur.next
			return true
		}
		prev = cur
		cur = cur.next
	}
	return false
}

// 反转链表
func (l *SingleLinkedList) Reverse() {
	l.Head = reverse(l.Head)
}

func reverse(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	newHead := reverse(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}

//打印链表
func (l *SingleLinkedList) Print() {
	cur := l.Head
	for cur != nil {
		fmt.Printf("[%d]", cur.data)
		cur = cur.next
	}
	fmt.Println("\n============")
}

// 移除对应值的节点
func (l *SingleLinkedList) DeleteAll(node *node) {
	cur := l.Head
	for cur != nil {
		l.Delete(node)

	}
}
