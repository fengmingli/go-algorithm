/**
 * @Author LFM
 * @Date 2021/9/6 8:30 下午
 * @Since V1
 */

package dp

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	paths := UniquePaths(20, 20)
	fmt.Println(paths)
}
