//930. 和相同的二元子数组  显示英文描述
//用户通过次数 88
//用户尝试次数 179
//通过次数 89
//提交次数 527
//题目难度 Medium
//在由若干 0 和 1  组成的数组 A 中，有多少个和为 S 的非空子数组。
//
//
//
//示例：
//
//输入：A = [1,0,1,0,1], S = 2
//输出：4
//解释：
//如下面黑体所示，有 4 个满足题目要求的子数组：
//[1,0,1,0,1]
//[1,0,1,0,1]
//[1,0,1,0,1]
//[1,0,1,0,1]
//
//
//提示：
//
//A.length <= 30000
//0 <= S <= A.length
//A[i] 为 0 或 1

package main

import "fmt"

func main() {
	A := []int{1, 0, 1, 0, 1}
	S := 2
	fmt.Println(numSubarraysWithSum(A, S))
}

func numSubarraysWithSum(A []int, S int) int {
	l := len(A)
	count := 0
	for i := 0; i < l; i++ {
		sum := A[i]
		if sum == S {
			count++
		}
		for j := i + 1; j < l; j++ {
			sum = sum + A[j]
			if sum > S {
				break
			} else if sum == S {
				count++
			}
		}
	}
	return count
}
