/**
 * @Author LFM
 * @Date 2021/8/15 9:44 下午
 * @Since V1
 */

package dp

//求最小值
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
