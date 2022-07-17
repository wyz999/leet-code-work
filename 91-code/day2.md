题目描述


```
给定一个字符串 S 和一个字符 C。返回一个代表字符串 S 中每个字符到字符串 S 中的字符 C 的最短距离的数组。

示例 1:

输入: S = "loveleetcode", C = 'e'
输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
说明:

- 字符串 S 的长度范围为 [1, 10000]。
- C 是一个单字符，且保证是字符串 S 里的字符。
- S 和 C 中的所有字母均为小写字母。
```


## 思路

- 先从左到右遍历，记录左侧到最近c的距离
- 从右到左，1，记录从右侧到最近c的距离,2，比较两个距离，取最小

## 代码

```
golang
func shortestToChar(s string, c byte) []int {
    n:=len(s)
    num :=make([]int, n)
    idx:=-n
    // 从左到右
  for k,v:=range s{
      if byte(v)==c{
        idx=k
      }
    num[k]=k-idx
  }
    // 从右到左
   idx=2*n
  for i:=n-1;i>=0;i--{
    if  s[i]== c{
      idx=i
    }
    num[i]=min(num[i],idx-i)
  }
  return num
}

func min(a,b int) int{
  if a>b {
    return b
  }
  return a
}
```

## 时空复杂度

- 时间复杂度：O(N)，其中 N 为数组长度。

* 空间复杂度：O(1)
