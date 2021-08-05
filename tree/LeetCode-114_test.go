/**
 * @Author LFM
 * @Date 2021/8/4 10:14 下午
 * @Since V1
 */

package tree

import (
	"fmt"
	"testing"
)

func TestBreadthTravel(t *testing.T) {
	root := perfectTree()
	root.PrintBT()
	fmt.Println()
	//Flatten(root)
	Falcon2(root)
	root.PrintBT()
	fmt.Println()

}
