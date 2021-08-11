/**
 * 双链表队列
 * <P>
 * 特性：先进先出
 * @Author LFM
 * @Date 2021/8/11 9:36 下午
 * @Since V1
 */

package queue

type (
	//Queue 队列
	Queue struct {
		top    *node
		rear   *node
		length int
	}
	//双向链表节点
	node struct {
		pre   *node
		next  *node
		value interface{}
	}
)

// NewLinkedQueue 创建一个null的队列
func NewLinkedQueue() *Queue {
	return &Queue{nil, nil, 0}
}

// Len 获取队列长度
func (q *Queue) Len() int {
	return q.length
}

// IsEmpty 返回true队列是否为null
func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

// Peek 返回队列顶端元素
func (q *Queue) Peek() interface{} {
	if q.top == nil {
		return nil
	}
	return q.top.value
}

// Push 入队操作
func (q *Queue) Push(v interface{}) {
	n := &node{nil, nil, v}
	if q.length == 0 {
		q.top = n
		q.rear = q.top
	} else {
		n.pre = q.rear
		q.rear.next = n
		q.rear = n
	}
	q.length++
}

// Pop 出队操作
// 返回队列顶端元素，并删除该元素
func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.top
	if q.top.next == nil {
		q.top = nil
	} else {
		q.top = q.top.next
		q.top.pre.next = nil
		q.top.pre = nil
	}
	q.length--
	return n.value
}
