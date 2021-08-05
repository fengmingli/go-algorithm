/**
* @Author LFM
* @Date 2021/8/4 9:54 下午
* @Since V1
 */

package tree

import "go-algorithm/structures"

type Node = structures.TreeNode

//Flatten 114 https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
//我们再梳理一下，如何按题目要求把一棵树拉平成一条链表？很简单，以下流程：
//1、将 root 的左子树和右子树拉平。
//2、将 root 的右子树接到左子树下方，然后将整个左子树作为右子树。
func Flatten(root *Node) {
	// 构建一个新的链表，然后再把root指过去
	if root == nil {
		return
	}
	newRoot := new(Node)
	pre := newRoot
	var inorderTravel func(root *Node)
	inorderTravel = func(root *Node) {
		if root == nil {
			return
		}

		left, right := root.Left, root.Right

		// 清楚left、right，构建单链表形式
		root.Left, root.Right = nil, nil
		pre.Right = root
		pre = pre.Right

		inorderTravel(left)
		inorderTravel(right)
	}
	inorderTravel(root)
}
