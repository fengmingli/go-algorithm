/**
 * @Author LFM
 * @Date 2021/8/15 9:45 下午
 * @Since V1
 */

package dp

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	var answer = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(answer))
}
