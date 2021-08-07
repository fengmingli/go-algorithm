/**
 * @Author LFM
 * @Date 2021/8/6 11:12 下午
 * @Since V1
 */

package tree

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	arr := []int{3, 2, 1, 6, 0, 5}
	tree := ConstructMaximumBinaryTree(arr)
	fmt.Print(tree)
}
