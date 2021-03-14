package main

// 归并排序 n*log2n

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{5, 4, 2, 1, 3, 9, 8, 7, 6}))
}

func mergeSort(nums []int) []int {
	length := len(nums)
	if length == 1 {
		return nums
	}
	mid := length / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge_sort(left, right)
}

func merge_sort(left []int, right []int) []int {
	lenL := len(left)
	lenR := len(right)
	length := lenL + lenR
	nums := make([]int, length)
	iL := 0
	iR := 0
	for i := 0; i < length; i++ {
		if iL < lenL && iR < lenR {
			if left[iL] < right[iR] {
				nums[i] = left[iL]
				iL++
			} else {
				nums[i] = right[iR]
				iR++
			}
		} else if iL < lenL {
			nums[i] = left[iL]
			iL++
		} else if iR < lenR {
			nums[i] = right[iR]
			iR++
		}
	}
	return nums
}
