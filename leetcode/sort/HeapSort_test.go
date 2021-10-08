/**
 * @Author LFM
 * @Date 2021/9/30 5:42 下午
 * @Since V1
 */

package sort

import (
	"fmt"
	"testing"
)

func TestHeadSort(t *testing.T) {

	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(arr)
	fmt.Println(HeapSort(arr))

}
