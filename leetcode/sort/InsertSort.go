/**
 * @Author LFM
 * @Date 2021/9/30 4:46 下午
 * @Since V1
 */

package sort

//InsertSort 插入排序 直接插入排序的时间复杂度是O(N2)
func InsertSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
	return s
}
