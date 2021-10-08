/**
 * @Author LFM
 * @Date 2021/9/30 5:35 下午
 * @Since V1
 */

package sort

func SelectSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}

	for i := 0; i < len(values); i++ {
		min := i // 初始的最小值位置从0开始，依次向右
		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := len(values) - 1; j > i; j-- {
			if values[j] < values[min] {
				min = j
			}
		}
		// 把每次找出来的最小值与之前的最小值做交换
		values[i], values[min] = values[min], values[i]
	}

	return values
}
