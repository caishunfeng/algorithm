package main

import "fmt"

//1030. 距离顺序排列矩阵单元格  显示英文描述
//用户通过次数 152
//用户尝试次数 179
//通过次数 152
//提交次数 253
//题目难度 Easy
//给出 R 行 C 列的矩阵，其中的单元格的整数坐标为 (r, c)，满足 0 <= r < R 且 0 <= c < C。
//
//另外，我们在该矩阵中给出了一个坐标为 (r0, c0) 的单元格。
//
//返回矩阵中的所有单元格的坐标，并按到 (r0, c0) 的距离从最小到最大的顺序排，其中，两单元格(r1, c1) 和 (r2, c2) 之间的距离是曼哈顿距离，|r1 - r2| + |c1 - c2|。（你可以按任何满足此条件的顺序返回答案。）
//
//
//
//示例 1：
//
//输入：R = 1, C = 2, r0 = 0, c0 = 0
//输出：[[0,0],[0,1]]
//解释：从 (r0, c0) 到其他单元格的距离为：[0,1]
//示例 2：
//
//输入：R = 2, C = 2, r0 = 0, c0 = 1
//输出：[[0,1],[0,0],[1,1],[1,0]]
//解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2]
//[[0,1],[1,1],[0,0],[1,0]] 也会被视作正确答案。
//示例 3：
//
//输入：R = 2, C = 3, r0 = 1, c0 = 2
//输出：[[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
//解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2,2,3]
//其他满足题目要求的答案也会被视为正确，例如 [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]]。
//
//
//提示：
//
//1 <= R <= 100
//1 <= C <= 100
//0 <= r0 < R
//0 <= c0 < C

func main() {
	fmt.Println(allCellsDistOrder(2, 3, 1, 2))
}

type Node struct {
	X int
	Y int
	L int
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	nodes := []*Node{}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			nodes = append(nodes, &Node{X: i, Y: j, L: abs(i, r0) + abs(j, c0)})
		}
	}
	sortByDist(nodes)
	result := [][]int{}
	l := len(nodes)
	for i := 0; i < l; i++ {
		result = append(result, []int{nodes[i].X, nodes[i].Y})
	}
	return result
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	} else {
		return a - b
	}
}

func sortByDist(nodes []*Node) {
	l := len(nodes)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nodes[i].L > nodes[j].L {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
		}
	}
}
