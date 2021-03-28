package main

import "fmt"

// import "strconv"

/**
*给你一个字符串 word ，该字符串由数字和小写英文字母组成。

请你用空格替换每个不是数字的字符。例如，"a123bc34d8ef34" 将会变成 " 123  34 8  34" 。注意，剩下的这些整数间至少要用一个空格隔开："123"、"34"、"8" 和 "34" 。

返回对 word 完成替换后形成的 不同 整数的数目。

如果两个整数的 不含前导零 的十进制表示不同，则认为这两个整数也不同。

1 <= word.length <= 1000
word 由数字和小写英文字母组成

示例 1：

输入：word = "a123bc34d8ef34"
输出：3
解释：不同的整数有 "123"、"34" 和 "8" 。注意，"34" 只计数一次。
示例 2：

输入：word = "leet1234code234"
输出：2
示例 3：

输入：word = "a1b01c001"
输出：1
解释："1"、"01" 和 "001" 视为同一个整数的十进制表示，因为在比较十进制值时会忽略前导零的存在。
*/

func main() {
	// fmt.Println(numDifferentIntegers("a123bc34d8ef34"))
	// fmt.Println(numDifferentIntegers("leet1234code234"))
	// fmt.Println(numDifferentIntegers("00a001b020"))

	str := "abcdefghijklmnopqrstuvwxyz0()"
	i := 0
	for(i < len(str)) {
		fmt.Println(str[i])
		i++
	}
}

// 遍历word每个字符，判断如果是小写字母则将上一次的字符串转为数字记录在数组中
func numDifferentIntegers(word string) int {
	l := len(word)
	index := 0
	start := 0
	strMap := make(map[string]bool)
	for index < l {
		// if start == index && word[index] == 48 {
		// 	start++
		// }
		if word[index] >= 97 && word[index] <= 122 {
			if index > start {
				str := zeroPrefix(word[start:index])
				strMap[str] = true
				// fmt.Println(str)
				start = index + 1
			} else {
				start++
			}
		}
		index++
	}
	if index > start {
		str := zeroPrefix(word[start:index])
		strMap[str] = true
		// fmt.Println(str)
	}
	return len(strMap)
}

func zeroPrefix(str string) string {
	l := len(str)
	begin := 0
	for i := 0; i < l; i++ {
		if str[i] == 48 {
			begin++
		} else {
			break
		}
	}
	if begin < l {
		return str[begin:]
	} else {
		return "0"
	}
}
