/**
 * @Author LFM
 * @Date 2021/8/4 7:49 下午
 * @Since V1
 */

package tree

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//创建二叉树
	root := NewNode(nil, nil, nil)
	root.SetData("root")

	a := NewNode(nil, nil, nil)
	a.SetData("left")

	al := NewNode(nil, nil, nil)
	al.SetData(100)

	ar := NewNode(nil, nil, nil)
	ar.SetData(3.14)

	a.Left = al
	a.Right = ar

	b := NewNode(nil, nil, nil)
	b.SetData("right")

	root.Left = a
	root.Right = b

	root.PrintBT()
	fmt.Println()

	// 使用 Order 接口实现对二叉树的基本操作
	var it Order  // 定义接口类型
	it = root     // 将*Node类型变量赋值给接口，*Node实现了接口的所有方法
	it.PreOrder() // 先序遍历
	fmt.Println()
	it.InOrder() // 中序遍历
	fmt.Println()
	it.PostOrder() // 后序遍历
	fmt.Println()
	it.BreadthTravel() // 广度遍历
}
