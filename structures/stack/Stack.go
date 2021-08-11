/**
 * 栈
 * <P>
 * 特性：先进后出
 * @Author LFM
 * @Date 2021/8/11 8:58 下午
 * @Since V1
 */

package stack

import (
	"container/list"
	"sync"
)

type Stack struct {
	list *list.List
	lock *sync.RWMutex
}

//NewStack 创建一个空栈
func NewStack() *Stack {
	l := list.New()
	lock := &sync.RWMutex{}
	return &Stack{l, lock}
}

// Push 入栈操作
func (stack *Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(value)
}

// Pop 出栈操作
// 返回栈顶元素
func (stack *Stack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

// Peak 返回栈顶元素
func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

// Len 返回栈长度
func (stack *Stack) Len() int {
	return stack.list.Len()
}

// IsEmpty 判断栈是否为空
func (stack *Stack) IsEmpty() bool {
	return stack.list.Len() == 0
}
