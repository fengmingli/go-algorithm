/**
 * @Author LFM
 * @Date 2021/8/11 9:57 下午
 * @Since V1
 */

package stack

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//入队列
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	//出队列 1234
	for !stack.IsEmpty() {
		fmt.Print(stack.Pop())
	}

}
