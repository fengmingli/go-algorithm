/**
 * 二叉树
 * @Author LFM
 * @Date 2021/8/11 8:21 下午
 * @Since V1
 */

package structures

import (
	"fmt"
	q "go-algorithm/structures/queue"
	s "go-algorithm/structures/stack"
)

//BiTreeNode 二叉树树的每一个节点数据结构
type BiTreeNode struct {
	//树的数据
	Data interface{}

	//树的左节点
	Left *BiTreeNode

	//树的右节点
	Right *BiTreeNode
}

type BiTree interface {
	PreOrder(tree *BiTreeNode) /*先序遍历*/

	InOrder(tree *BiTreeNode) /*中序遍历*/

	PostOrder(tree *BiTreeNode) /*后续遍历*/

	BfOrder(tree *BiTreeNode) /*广度优先遍历*/

	DfOrder(tree *BiTreeNode) /*深度优先遍历*/
}

// PreOrder 先序遍历 root->left->right
func (n BiTreeNode) PreOrder(tree *BiTreeNode) {
	if tree == nil {
		return
	}
	fmt.Print(tree.Data, " ")
	n.PreOrder(tree.Left)
	n.PreOrder(tree.Right)
}

// InOrder 中序遍历 left->root->right
func (n BiTreeNode) InOrder(tree *BiTreeNode) {
	if tree == nil {
		return
	}
	n.InOrder(tree.Left)
	fmt.Print(tree.Data, " ")
	n.InOrder(tree.Right)
}

// PostOrder 后序遍历 left->right—>root
func (n BiTreeNode) PostOrder(tree *BiTreeNode) {
	if tree == nil {
		return
	}
	n.PostOrder(tree.Left)
	n.PostOrder(tree.Right)
	fmt.Print(tree.Data, " ")
}

//BfOrder 广度优先遍历：从根节点一层层从左向右遍历
func (n BiTreeNode) BfOrder(tree *BiTreeNode) {
	if tree == nil {
		fmt.Println("BfOrder : tree is empty")
		return
	}
	element := make([]interface{}, 0)
	queue := q.NewLinkedQueue()
	/*根节点先入队列*/
	queue.Push(tree)
	for !queue.IsEmpty() {
		/*出队列*/
		node := queue.Pop().(*BiTreeNode)
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
		element = append(element, node.Data)
	}
	for i := 0; i < len(element); i++ {
		fmt.Print(element[i], " ")
	}
	fmt.Println()
}

//DfOrder 深度优先：和前序遍历一样
func (n BiTreeNode) DfOrder(tree *BiTreeNode) {
	if tree == nil {
		fmt.Println("DfOrder : tree is empty")
		return
	}
	element := make([]interface{}, 0)
	stack := s.NewStack()
	stack.Push(&n)
	for !stack.IsEmpty() {
		node := stack.Pop().(*BiTreeNode)
		element = append(element, node.Data)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
	for i := 0; i < len(element); i++ {
		fmt.Print(element[i], " ")
	}
	fmt.Println()
}
