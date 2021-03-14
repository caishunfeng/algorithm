//931. 下降路径最小和  显示英文描述
//用户通过次数 72
//用户尝试次数 109
//通过次数 74
//提交次数 202
//题目难度 Medium
//给定一个方形整数数组 A，我们想要得到通过 A 的下降路径的最小和。
//
//下降路径可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列。
//
//
//
//示例：
//
//输入：[[1,2,3],[4,5,6],[7,8,9]]
//输出：12
//解释：
//可能的下降路径有：
//[1,4,7], [1,4,8], [1,5,7], [1,5,8], [1,5,9]
//[2,4,7], [2,4,8], [2,5,7], [2,5,8], [2,5,9], [2,6,8], [2,6,9]
//[3,5,7], [3,5,8], [3,5,9], [3,6,8], [3,6,9]
//和最小的下降路径是 [1,4,7]，所以答案是 12。
//
//
//
//提示：
//
//1 <= A.length == A[0].length <= 100
//-100 <= A[i][j] <= 100

package main

import (
	"fmt"
	"math"
)

func main() {
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(A))
}

func minFallingPathSum(A [][]int) int {
	minSum := 0
	lx := len(A)
	ly := len(A[0])

	sums := make([][]int, lx)

	for i := 0; i < lx; i++ {
		sums[i] = make([]int, ly)
	}

	for j := 0; j < ly; j++ {
		sums[0][j] = A[0][j]
	}

	for i := 1; i < lx; i++ {
		for j := 0; j < ly; j++ {
			min := math.MaxInt64
			if j-1 >= 0 && sums[i-1][j-1] < min {
				min = sums[i-1][j-1]
			}
			if sums[i-1][j] < min {
				min = sums[i-1][j]
			}
			if j+1 < ly && sums[i-1][j+1] < min {
				min = sums[i-1][j+1]
			}
			sums[i][j] = min + A[i][j]
		}
	}

	//fmt.Println(sums)

	minSum = sums[lx-1][0]
	for j := 1; j < ly; j++ {
		if minSum > sums[lx-1][j] {
			minSum = sums[lx-1][j]
		}
	}

	return minSum
}
