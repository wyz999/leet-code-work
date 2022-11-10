/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	n := len(digits)
	for i:=n-1;i>=0;i--{
		if digits[i]==9{
			digits[i]=0
		}else{
			digits[i]+=1
			return digits
		}
	}
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
// @lc code=end

