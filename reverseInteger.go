package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d\n", reverse(1534236469))
}

func reverse(x int) int {
	var y int
	min := -math.Pow(2, 31)
	max := math.Pow(2, 31) - 1
	for {
		if x == 0 {
			break
		}
		y = y*10 + x%10
		x = x / 10
		if float64(y) < min || float64(y) > max {
			return 0
		}
	}
	return y
}

//Container With Most Water
func maxArea(height []int) int {
	var maxV, currentV int
	for i := 0; i < len(height); i++ {
		for j := i; j < len(height); j++ {
			currentV = (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
			if currentV > maxV {
				maxV = currentV
			}
		}
	}
	return maxV
}
