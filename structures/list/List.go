/**
 * @Author LFM
 * @Date 2021/8/16 3:12 下午
 * @Since V1
 */

package list

import "sync"

type Data interface{}

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type List struct {
	mutex sync.RWMutex
	len   int
	head  *Node
	tail  *Node
}

func NewNode(data interface{}) (node *Node) {
	node = &Node{
		data: data,
		prev: nil,
		next: nil,
	}
	return
}

func (n *Node) GetDataPointer() (d *interface{}) {
	d = &n.data
	return
}

func (n *Node) GetData() (d interface{}) {
	d = n.data
	return
}

func NewList() (list *List) {
	list = &List{}
	return
}

func (l *List) RPush(node *Node) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.len == 0 {
		l.head = node
		l.tail = node
	} else {
		tail := l.tail
		tail.next = node
		node.prev = tail

		l.tail = node
	}

	l.len = l.len + 1
}

func (l *List) LPop() (node *Node) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.len > 0 {
		node = l.head
		if node.next == nil {
			l.head = nil
			l.tail = nil
		} else {
			l.head = node.next
		}
		l.len = l.len - 1
	}
	return
}

func (l *List) LPopAll() (nodes []*Node) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	for l.len > 0 {
		var node = l.head
		if node.next == nil {
			l.head = nil
			l.tail = nil
		} else {
			l.head = node.next
		}
		l.len = l.len - 1
		nodes = append(nodes, node)
	}
	return
}

func (l *List) Len() (len int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	len = l.len
	return
}
