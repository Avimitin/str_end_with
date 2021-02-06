package main

import "fmt"

func solution(str, ending string) bool {
	strLen := len(str)
	endLen := len(ending)
	if endLen > strLen {
		return false
	}
	var i, j = strLen - 1, endLen - 1
	for j >= 0 {
		if ending[j] == ' ' {
			j--
			continue
		}
		if str[i] != ending[j] {
			return false
		}
		i--
		j--
	}
	return true
}

func main() {
	fmt.Println(solution("", ""))
	fmt.Println(solution(" ", ""))
	fmt.Println(solution("abc", "c"))
	fmt.Println(solution("banana", "ana"))
	fmt.Println(solution("a", "z"))
	fmt.Println(solution("", "t"))
	fmt.Println(solution("$a = $b + 1", "+1"))
	fmt.Println(solution("    ", "   "))
	fmt.Println(solution("1oo", "100"))
}
