/**
 * @Author LFM
 * @Date 2021/9/30 4:33 下午
 * @Since V1
 */

package sort

//冒泡排序 核心思想：相邻比较，大的替换小的，每次排列都能冒出一个最大值
//时间复杂度：
//遍历一趟的时间复杂度是O(N)，需要遍历多少次呢?
//N-1次！因此，冒泡排序的时间复杂度是O(N2)
func bubbleSort(source []int) []int {
	if len(source) <= 1 {
		return source
	}

	for i := 0; i < len(source)-1; i++ {
		for j := i + 1; j < len(source); j++ {
			if source[i] > source[j] {
				source[j], source[i] = source[i], source[j]
			}
		}
	}
	return source
}
