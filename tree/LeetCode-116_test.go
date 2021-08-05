/**
 * @Author LFM
 * @Date 2021/8/4 9:33 下午
 * @Since V1
 */

package tree

import (
	"fmt"
	"testing"
)

func perfectTree() *Node {
	//创建二叉树
	root := NewNode(nil, nil, nil)
	root.SetData(1)

	l21 := NewNode(nil, nil, nil)
	l21.SetData(2)

	l22 := NewNode(nil, nil, nil)
	l22.SetData(3)

	l31 := NewNode(nil, nil, nil)
	l31.SetData(4)

	l32 := NewNode(nil, nil, nil)
	l32.SetData(5)

	l33 := NewNode(nil, nil, nil)
	l33.SetData(6)

	l34 := NewNode(nil, nil, nil)
	l34.SetData(7)

	temp := NewNode(nil, nil, nil)
	temp.SetData("#")

	root.Left = l21
	root.Right = l22
	root.Next = temp

	l21.Left = l31
	l21.Right = l32
	l21.Next = l22

	//l22.Left = l33
	l22.Right = l34
	l22.Next = temp

	l31.Next = l32
	l32.Next = l33
	l33.Next = l34
	l34.Next = temp

	return root
}

func TestConnect(t *testing.T) {
	root := perfectTree()

	root.PrintBT()
	fmt.Println()
	connect(root)

	root.PrintBT()
	fmt.Println()

}
