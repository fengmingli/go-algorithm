/**
 * @Author LFM
 * @Date 2021/8/4 9:54 下午
 * @Since V1
 */

package tree

func flatten(root *Node) {
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
