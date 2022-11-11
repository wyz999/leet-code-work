/*
 * @lc app=leetcode.cn id=989 lang=golang
 *
 * [989] 数组形式的整数加法
 num = [2,1,5], k = 806
 */

// @lc code=start
func addToArrayForm(num []int, k int) []int {
	var res []int
	for i := len(num)-1;i >= 0;i--{
		// 当前位
		sum := num[i] + k%10
		k /= 10
		// 当前位大于10
		if sum >= 10{
			//进位
			k++ 
			sum -= 10
		}
		res = append(res,sum)
	}
	for ;k>0;k/=10{
		res = append(res,k%10)
	}
	reverse(res)
	return res
}

func reverse(nums []int){
	for i,j := 0,len(nums);i<j/2;i++{
		nums[i],nums[j-1-i] = nums[j-1-i],nums[i]
	}
}
// @lc code=end

