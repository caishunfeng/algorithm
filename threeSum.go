package main

import "fmt"

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	if n == 0 || nums[0] > 0 || nums[n-1] < 0 {
		return result
	}

	target := 0
	i := 0
	j := 0
	for k := 0; k < n; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		target = 0 - nums[k]
		i = k + 1
		j = n - 1
		for i < j {
			if nums[i]+nums[j] == target {
				result = append(result, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else {
				j--
			}
		}

	}
	return result
}
