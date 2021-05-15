package main

import "fmt"

// 力扣团队买了一个可编程机器人，机器人初始位置在原点(0, 0)。小伙伴事先给机器人输入一串指令command，机器人就会无限循环这条指令的步骤进行移动。指令有两种：

// U: 向y轴正方向移动一格
// R: 向x轴正方向移动一格。
// 不幸的是，在 xy 平面上还有一些障碍物，他们的坐标用obstacles表示。机器人一旦碰到障碍物就会被损毁。

// 给定终点坐标(x, y)，返回机器人能否完好地到达终点。如果能，返回true；否则返回false。

// 示例 1：

// 输入：command = "URR", obstacles = [], x = 3, y = 2
// 输出：true
// 解释：U(0, 1) -> R(1, 1) -> R(2, 1) -> U(2, 2) -> R(3, 2)。
// 示例 2：

// 输入：command = "URR", obstacles = [[2, 2]], x = 3, y = 2
// 输出：false
// 解释：机器人在到达终点前会碰到(2, 2)的障碍物。
// 示例 3：

// 输入：command = "URR", obstacles = [[4, 2]], x = 3, y = 2
// 输出：true
// 解释：到达终点后，再碰到障碍物也不影响返回结果。
//

// 限制：

// 2 <= command的长度 <= 1000
// command由U，R构成，且至少有一个U，至少有一个R
// 0 <= x <= 1e9, 0 <= y <= 1e9
// 0 <= obstacles的长度 <= 1000
// obstacles[i]不为原点或者终点

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/programmable-robot
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	// fmt.Println(robot("URR", [][]int{}, 3, 2))
	// fmt.Println(robot("URR", [][]int{{2, 2}}, 3, 2))
	// fmt.Println(robot("URR", [][]int{{4, 2}}, 3, 2))
	// fmt.Println(robot("RRU", [][]int{{5, 5}, {9, 4}, {9, 7}, {6, 4}, {7, 0}, {9, 5}, {10, 7}, {1, 1}, {7, 5}}, 1486, 743))
	// fmt.Println(robot("URRUU", [][]int{{6, 0}, {3, 4}, {5, 2}, {3, 0}}, 6, 6))
	// fmt.Println(robot("RUURRR", [][]int{{9, 5}, {4, 4}, {0, 10}, {3, 9}}, 80, 40))
	fmt.Println(robot("UUUURR", [][]int{{0, 8}, {1, 2}, {7, 3}, {1, 5}, {7, 9}, {4, 9}, {3, 3}}, 990, 1890))
}

func robot(command string, obstacles [][]int, x int, y int) bool {
	baseX := 0
	baseY := 0

	for i := 0; i < len(obstacles)-1; i++ {
		for j := i + 1; j < len(obstacles); j++ {
			if obstacles[i][0] > obstacles[j][0] {
				obstacles[i], obstacles[j] = obstacles[j], obstacles[i]
			} else if obstacles[i][0] == obstacles[j][0] && obstacles[i][1] > obstacles[j][1] {
				obstacles[i], obstacles[j] = obstacles[j], obstacles[i]
			}
		}
	}

	for true {
		for i := 0; i < len(command); i++ {
			switch command[i] {
			case 'U':
				baseY += 1
				break
			case 'R':
				baseX += 1
				break
			default:
				return false
			}
			if isMeet(obstacles, baseX, baseY) {
				return false
			}
			if baseX == x && baseY == y {
				return true
			}
			if baseX > x || baseY > y {
				return false
			}
		}
	}
	return false
}

func isMeet(obstacles [][]int, baseX int, baseY int) bool {
	l := len(obstacles)
	if l == 0 {
		return false
	}
	if l == 1 {
		if obstacles[l-1][0] == baseX && obstacles[l-1][1] == baseY {
			return true
		} else {
			return false
		}
	}

	mid := 0
	if l%2 == 0 {
		mid = l / 2
	} else {
		mid = l/2 + 1
	}
	if obstacles[mid-1][0] < baseX || obstacles[mid-1][1] < baseY {
		return isMeet(obstacles[mid:], baseX, baseY)
	} else {
		return isMeet(obstacles[:mid], baseX, baseY)
	}
}
