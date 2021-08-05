package tree

// InvertTree 力扣 226 https://leetcode-cn.com/problems/invert-binary-tree/
func InvertTree(root *Node) *Node {
	// base case
	if root == nil {
		return nil
	}
	/**** 前序/后序遍历位置都可以 ****/
	// root 节点需要交换它的左右子节点
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	// 让左右子节点继续翻转它们的子节点
	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}
