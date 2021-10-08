/**
 * @Author LFM
 * @Date 2021/9/23 10:42 下午
 * @Since V1
 */

package _6_Jump_Gamp

func jumpGame(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	maxJump := 0
	for k, v := range nums {
		if k > maxJump {
			return false
		}
		maxJump = max(maxJump, k+v)
	}
	return true
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
