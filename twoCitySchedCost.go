package main

import (
	"fmt"
)

//1029. 两地调度  显示英文描述
//用户通过次数 0
//用户尝试次数 0
//通过次数 0
//提交次数 0
//题目难度 Easy
//公司计划面试 2N 人。第 i 人飞往 A 市的费用为 costs[i][0]，飞往 B 市的费用为 costs[i][1]。
//
//返回将每个人都飞到某座城市的最低费用，要求每个城市都有 N 人抵达。
//
//
//
//示例：
//
//输入：[[10,20],[30,200],[400,50],[30,20]]
//输出：110
//解释：
//第一个人去 A 市，费用为 10。
//第二个人去 A 市，费用为 30。
//第三个人去 B 市，费用为 50。
//第四个人去 B 市，费用为 20。
//
//最低总费用为 10 + 30 + 50 + 20 = 110，每个城市都有一半的人在面试。
//
//
//提示：
//
//1 <= costs.length <= 100
//costs.length 为偶数
//1 <= costs[i][0], costs[i][1] <= 1000

func main() {
	costs := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	fmt.Println(twoCitySchedCost(costs))
}

type CostIndex struct {
	Index    int
	Diff     int
	MinIndex int
	Min      int
}

func twoCitySchedCost(costs [][]int) int {
	minTotal := 0
	l := len(costs)
	n := l / 2
	costIndexs := []*CostIndex{}
	for i := 0; i < l; i++ {
		costIndex := &CostIndex{
			Index:    i,
			Diff:     getAbs(costs[i][0] - costs[i][1]),
			MinIndex: getMinIndex(costs[i][0], costs[i][1]),
			Min:      getMin(costs[i][0], costs[i][1]),
		}
		costIndexs = append(costIndexs, costIndex)
	}
	sort(costIndexs)

	//for i := 0; i < l; i++ {
	//	fmt.Printf("%+v\n", costIndexs[i])
	//}

	countA := 0
	countB := 0
	for i := 0; i < l; i++ {
		//fmt.Printf("i:%d,total:%d,countA:%d,countB:%d\n", i, minTotal, countA, countB)
		costIndex := costIndexs[i]
		if costIndex.MinIndex == 0 && countA < n {
			minTotal = minTotal + costs[costIndex.Index][0]
			countA++
		} else if costIndex.MinIndex == 1 && countB < n {
			minTotal = minTotal + costs[costIndex.Index][1]
			countB++
		} else if costIndex.MinIndex == 0 && countA == n {
			minTotal = minTotal + costs[costIndex.Index][1]
			countB++
		} else if costIndex.MinIndex == 1 && countB == n {
			minTotal = minTotal + costs[costIndex.Index][0]
			countA++
		}
	}

	return minTotal
}

func sort(costIndexs []*CostIndex) {
	l := len(costIndexs)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if costIndexs[i].Diff < costIndexs[j].Diff {
				costIndexs[i], costIndexs[j] = costIndexs[j], costIndexs[i]
			} else if costIndexs[i].Diff == costIndexs[j].Diff {
				if costIndexs[i].Min > costIndexs[j].Min {
					costIndexs[i], costIndexs[j] = costIndexs[j], costIndexs[i]
				}
			}
		}
	}
}

func getAbs(diff int) int {
	if diff < 0 {
		return -diff
	} else {
		return diff
	}
}

func getMinIndex(a, b int) int {
	if a > b {
		return 1
	} else {
		return 0
	}
}

func getMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
