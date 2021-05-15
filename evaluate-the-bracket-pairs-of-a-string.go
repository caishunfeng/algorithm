package main

import "fmt"

func main() {
	fmt.Println(evaluate("(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}}))
}

func evaluate(s string, knowledge [][]string) string {
	kMap := make(map[string]string)
	for _, k := range knowledge {
		kMap[k[0]] = k[1]
	}

	result := ""
	start := -1
	end := -1
	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			start = i
		} else if s[i] == 41 {
			end = i
			key := s[start+1 : end]
			// fmt.Println(key)
			if kMap[key] != "" {
				result = result + kMap[key]
			} else {
				result = result + "?"
			}
			start = -1
		} else if start == -1 {
			result = result + s[i:i+1]
		}
	}
	return result
}
