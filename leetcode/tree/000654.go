/**
 * @Author LFM
 * @Date 2021/8/6 10:44 下午
 * @Since V1
 */

package tree

// ConstructMaximumBinaryTree https://leetcode-cn.com/problems/maximum-binary-tree/
// 思路：1、对于根节点 ，即找到nums中最大的值与对应索引；
//      2、分别递归，左右数组构建左右子树
func ConstructMaximumBinaryTree(nums []int) *Node {
	if len(nums) < 1 {
		return nil
	}
	//首选找到最大值
	index := findMax(nums)
	//其次构造二叉树
	root := &Node{
		Val:   nums[index],
		Left:  ConstructMaximumBinaryTree(nums[:index]),   //左节点
		Right: ConstructMaximumBinaryTree(nums[index+1:]), //右节点
	}
	return root
}

func findMax(nums []int) (index int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return
}
