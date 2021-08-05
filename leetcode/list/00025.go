/**
* @Author LFM
* @Date 2021/8/5 10:36 下午
* @Since V1
 */

package list

import (
	"go-algorithm/structures"
)

// Node ListNode define
type Node = structures.ListNode

//reverseKGroup  https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

func reverseKGroup(head *Node) *Node {
	dummy := &Node{Next: head}
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
}
