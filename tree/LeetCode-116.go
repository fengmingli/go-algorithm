/**
 * @Author LFM
 * @Date 2021/8/4 9:24 下午
 * @Since V1
 */

package tree

func connect(root *Node) {

}

//Connect 力扣 116 https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmp(root.Left, root.Right)
	return root
}

func tmp(left, right *Node) {
	if left == nil || right == nil {
		return
	}
	left.Next = right
	tmp(left.Left, left.Right)
	tmp(right.Left, right.Right)
	tmp(left.Right, right.Left)
}
