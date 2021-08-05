package tree

import (
	"fmt"
	"testing"
)

func initBinaryTree() *Node {
	//创建二叉树
	root := NewNode(nil, nil, nil)
	root.SetData(4)

	l21 := NewNode(nil, nil, nil)
	l21.SetData(2)

	l22 := NewNode(nil, nil, nil)
	l22.SetData(7)

	l31 := NewNode(nil, nil, nil)
	l31.SetData(1)

	l32 := NewNode(nil, nil, nil)
	l32.SetData(3)

	l33 := NewNode(nil, nil, nil)
	l33.SetData(6)

	l34 := NewNode(nil, nil, nil)
	l34.SetData(9)

	root.Left = l21
	root.Right = l22

	l21.Left = l31
	l21.Right = l32

	l22.Left = l33
	l22.Right = l34

	return root
}

func TestInvertTree(t *testing.T) {
	root := initBinaryTree()

	root.PrintBT()
	fmt.Println()

	InvertTree(root)

	root.PrintBT()
	fmt.Println()

}
