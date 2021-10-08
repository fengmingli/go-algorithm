/**
 * @Author LFM
 * @Date 2021/9/30 4:12 下午
 * @Since V1
 */

package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	s := []int{6, 7, 8, 10, 4, 6, 99}
	res := mergeSort(s)
	fmt.Println(res)
}
