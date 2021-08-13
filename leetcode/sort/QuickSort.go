/**
 * @Author LFM
 * @Date 2021/8/13 11:17 下午
 * @Since V1
 */

package sort

func quickSort(sortArray []int, left, right int) {
	if left < right {
		key := sortArray[left]
		i := left
		j := right
		for {
			for sortArray[i] < key {
				i++
			}
			for sortArray[j] > key {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}

		quickSort(sortArray, left, i-1)
		quickSort(sortArray, j+1, right)
	}
}
