package main

import "fmt"

//["","","","","(())()()","","((())())","()()(())","(()(()))","","()((()))",""]
//["","","((())())","","(()(()))","","","(())(())","(())()()","()((()))","","","()()(())"]

func main() {
	//strsA := []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}
	//strsB := []string{"()()()()", "(()()())", "(()())()", "()(()())", "((()()))", "(())()()", "()(())()", "((())())", "()()(())", "(()(()))", "((()))()", "()((()))", "(((())))"}
	//fmt.Println(findNotMatch(strsA, strsB))
	fmt.Println(len(generateParenthesis(4)))
}

func generateParenthesis(n int) []string {
	result := []string{}
	if n < 1 {
		return result
	}
	if n == 1 {
		result = append(result, "()")
		return result
	}
	mmap := make(map[string]bool)
	childs := generateParenthesis(n - 1)
	for _, child := range childs {
		for index := 0; index < len(child); index++ {
			str := child[:index] + "()" + child[index:]
			if !mmap[str] {
				result = append(result, str)
				mmap[str] = true
			}
		}
	}
	return result
}

func isValidParenthesis(nums []int) bool {
	sum := 0
	for _, num := range nums {
		if num+sum > 0 {
			return false
		} else {
			sum = sum + num
		}
	}
	return true
}

func printParenthesis(nums []int) string {
	var result string
	for _, num := range nums {
		if num == -1 {
			result = result + "("
		} else {
			result = result + ")"
		}
	}
	return result
}

func findNotMatch(strsA []string, strsB []string) []string {
	result := []string{}
	if len(strsA) < len(strsB) {
		strsA, strsB = strsB, strsA
	}

	for i := 0; i < len(strsA); i++ {
		isMatch := false
		for j := 0; j < len(strsB); j++ {
			if strsA[i] == strsB[j] {
				isMatch = true
				break
			}
		}
		if !isMatch {
			result = append(result, strsA[i])
		}
	}
	return result
}
