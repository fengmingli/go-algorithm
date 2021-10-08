/**
 * @Author LFM
 * @Date 2021/9/30 5:36 下午
 * @Since V1
 */

package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	numbers := []int{6, 2, 7, 5, 8, 9}
	fmt.Println(numbers)
	sort := SelectSort(numbers)
	fmt.Println(sort)
}
