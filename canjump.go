package main

import "fmt"

func main() {
	//fmt.Println(jump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	n := len(nums)

	if n == 0 {
		return false
	}

	if nums[0] >= n-1 {
		return true
	}

	list := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		if nums[i] >= n-i-1 {
			list[i] = 1
		} else {
			for j := n - 1; j > i; j-- {
				if nums[i] >= (j-i) && list[j] == 1 {
					list[i] = 1
					break
				}
			}
		}
	}

	fmt.Println(list)

	if list[0] == 1 {
		return true
	} else {
		return false
	}
}

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[i] = int(^uint(0) >> 1)
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] == 0 {
			continue
		}
		if nums[i] >= n-1-i {
			count[i] = 1
		} else {
			min := int(^uint(0) >> 1)
			for j := i + 1; j <= i+nums[i]; j++ {
				if count[j] < min {
					min = count[j]
				}
			}
			if min != int(^uint(0)>>1) {
				count[i] = min + 1
			}
		}
	}
	return count[0]
}
