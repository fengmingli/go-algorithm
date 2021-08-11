/**
 * @Author LFM
 * @Date 2021/8/11 9:50 下午
 * @Since V1
 */

package queue

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//入队列
	queue := NewLinkedQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	//出队列 1234
	for !queue.IsEmpty() {
		fmt.Print(queue.Pop())
	}

}
