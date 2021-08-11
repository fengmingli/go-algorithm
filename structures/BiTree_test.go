/**
 * @Author LFM
 * @Date 2021/8/11 8:34 下午
 * @Since V1
 */

package structures

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	tree := &BiTreeNode{
		Data: "A",
		Left: &BiTreeNode{
			Data: "B",
			Left: &BiTreeNode{
				Data:  "D",
				Left:  nil,
				Right: nil,
			},
			Right: &BiTreeNode{
				Data: "C",
				Left: nil,
				Right: &BiTreeNode{
					Data:  "H",
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &BiTreeNode{
			Data: "E",
			Left: &BiTreeNode{
				Data:  "G",
				Left:  nil,
				Right: nil,
			},
			Right: &BiTreeNode{
				Data:  "F",
				Left:  nil,
				Right: nil,
			},
		},
	}

	/*前序遍历*/
	tree.PreOrder(tree)
	fmt.Println()
	/*中序遍历*/
	tree.InOrder(tree)
	fmt.Println()
	/*后续遍历*/
	tree.PostOrder(tree)
	fmt.Println()
	/*广度优先*/
	tree.BfOrder(tree)
	/*深度优先*/
	tree.DfOrder(tree)
}
