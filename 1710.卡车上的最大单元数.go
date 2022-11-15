/*
 * @lc app=leetcode.cn id=1710 lang=golang
 *
 * [1710] 卡车上的最大单元数
 */

// @lc code=start
func maximumUnits(boxTypes [][]int, truckSize int) int {
	l := len(boxTypes)
	ans := 0
	for i:=0;i<l;i++{
			for j:=i+1;j<l;j++{
					if boxTypes[j][1]>boxTypes[i][1]{
							boxTypes[i],boxTypes[j] = boxTypes[j],boxTypes[i]
					}
			} 
	}
	
	for i:=0;i<l&&truckSize>0;i++{
			ans += boxTypes[i][0] * boxTypes[i][1]
			truckSize-=boxTypes[i][0]
			if truckSize < 0 {
		ans += boxTypes[i][1] * truckSize
	}
	}
	return ans
}
// @lc code=end

