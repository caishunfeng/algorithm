package main

import "fmt"

func main() {
	//nums := []int{1, 0, -1, 0, -2, 2}
	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
	fmt.Println(fourSum(nums, -11))
}

func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	result := [][]int{}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	threeSum := 0
	for i := l - 1; i >= 3; i-- {
		if i < l-1 && nums[i] == nums[i+1] {
			continue
		}
		threeSum = target - nums[i]
		threeSumResults := threeTarget(nums[:i], threeSum)
		for _, threeSumResult := range threeSumResults {
			result = append(result, append(threeSumResult, nums[i]))
		}
	}
	return result
}

func threeTarget(nums []int, target int) [][]int {
	fmt.Println(nums, target)
	l := len(nums)
	result := [][]int{}

	twoSum := 0
	i := 0
	j := 0
	for k := 0; k < l; k++ {
		twoSum = target - nums[k]
		//if nums[k] > twoSum {
		//	break
		//}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i = k + 1
		j = l - 1
		for i < j {
			if nums[i]+nums[j] == twoSum {
				result = append(result, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < twoSum {
				i++
			} else {
				j--
			}
		}

	}
	return result
}
