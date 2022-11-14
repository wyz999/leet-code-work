/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	end := len(s) -1
	ans := 0
	for s[end] == ' '{
			end--
	}
	for end>=0&&s[end]!=' '{
			end--
			ans++
	}
	return ans

}
// @lc code=end

