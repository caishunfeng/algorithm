package main

import "fmt"

// 600. 不含连续1的非负整数
// 给定一个正整数 n，找出小于或等于 n 的非负整数中，其二进制表示不包含 连续的1 的个数。

// 示例 1:

// 输入: 5
// 输出: 5
// 解释:
// 下面是带有相应二进制表示的非负整数<= 5：
// 0 : 0
// 1 : 1
// 2 : 10
// 3 : 11
// 4 : 100
// 5 : 101
// 6 : 110
// 7 : 111
// 8 : 1000
// 9 : 1001
// 10 : 1010
// 11 : 1011
// 其中，只有整数3违反规则（有两个连续的1），其他5个满足规则。
// 说明: 1 <= n <= 109

func main() {
	fmt.Println(findIntegers(2))
	fmt.Println(findIntegers(3))
	fmt.Println(findIntegers(4))
	fmt.Println(findIntegers(5))
	fmt.Println(findIntegers(6))
	fmt.Println(findIntegers(10))
}

var numMap = make(map[int]int)

func findIntegers(n int) int {
	cnt := 0
	k := 2
	round := 1
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 3
	}
	for n > k {
		if val, ok := numMap[round]; ok && val > 0 {
			continue
		} else {
			numMap[round] = findIntegers(k)
		}
		k = k * 2
		round++
	}
	cnt = numMap[round-1] + findIntegers(n-k)
	return cnt
}
