/*
 * @lc app=leetcode.cn id=1704 lang=golang
 *
 * [1704] 判断字符串的两半是否相似
 */

// @lc code=start
func halvesAreAlike(s string) bool {
	cun := 0
	for _,v := range s[:len(s)/2]{
			if strings.ContainsRune("aeiouAEIOU",v){
					cun++
			}
	}
	 for _,v := range s[len(s)/2:]{
			if strings.ContainsRune("aeiouAEIOU",v){
					cun--
			}
	}
	return cun==0
}
// @lc code=end

