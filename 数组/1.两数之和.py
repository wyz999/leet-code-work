#
# @lc app=leetcode.cn id=1 lang=python
#
# [1] 两数之和
#

# @lc code=start
class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for x in range(len(nums)):
            z = target-nums[x]
            if z in nums:
                y = nums.index[z]
                if(z!=y):
                    return x,y
            

# @lc code=end

