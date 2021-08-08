/**
 * @Author LFM
 * @Date 2021/8/7 3:30 下午
 * @Since V1
 */

package tree

import (
	"fmt"
)

//解题思路 中序
//利用递归遍历的特性从节点构造出一个序列化字符串
//1、将二叉树通过中序遍历方式，序列化
//2、根据序列化结果，建立哈希表
//3、统计每次序列化的结果出现的次数，一旦发现出现两次了，加入结果集

func findDuplicateSubtrees(root *Node) []*Node {
	if root == nil {
		return nil
	}
	//创建一个map用来保存 序列化化过得root左右节点 root，left，right
	m := make(map[string]int)
	//保存返回的结果
	var result []*Node
	var dfs func(root *Node) string

	dfs = func(root *Node) string {
		if root == nil {
			return ""
		}
		serial := fmt.Sprintf("%d,%s,%s", root.Val, dfs(root.Left), dfs(root.Right))
		m[serial]++
		if m[serial] == 2 {
			result = append(result, root)
		}
		return serial
	}
	dfs(root)
	return result
}
