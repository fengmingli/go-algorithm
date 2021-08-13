/**
 * @Author LFM
 * @Date 2021/8/13 11:36 下午
 * @Since V1
 */

package sort

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	arr := []int{5, 3, 2, 1, 7, 6, 9, 8, 4}
	quickSort(arr, 0, len(arr)-1)

	for i := 0; i < len(arr)-1; i++ {
		fmt.Print(arr[i], " ")
	}
}
