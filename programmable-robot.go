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
	fmt.Println(robot("URR", [][]int{}, 3, 2))
	fmt.Println(robot("URR", [][]int{{2, 2}}, 3, 2))
	fmt.Println(robot("URR", [][]int{{4, 2}}, 3, 2))
	fmt.Println(robot("URR", [][]int{}, 0, 0))
}

func robot(command string, obstacles [][]int, x int, y int) bool {
	// if len(command) < 2 || len(command) > 1000 {
	// 	return false
	// }

	// if !strings.Contains(command, "U") || !strings.Contains(command, "R") {
	// 	return false
	// }

	// if x < 0 || x > 1e9 {
	// 	return false
	// }

	// if y < 0 || y > 1e9 {
	// 	return false
	// }

	// if len(obstacles) > 1000 {
	// 	return false
	// }
	// var sites [][]int

	// fmt.Println(obstacles)
	// for i := 0; i < len(obstacles); i++ {
	// 	for j := 0; j < len(obstacles[i]); j++ {
	// 		fmt.Println(obstacles[i][j])
	// 	}
	// }

	baseX := 0
	baseY := 0

	for true {
		for i := 0; i < len(command); i++ {
			// fmt.Println(baseX, baseY)
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
			for k := 0; k < len(obstacles); k++ {
				if obstacles[k][0] == baseX && obstacles[k][1] == baseY {
					return false
				}
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
