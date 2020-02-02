package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	chkr := ""
	currStart := 0

	for index, c := range s {
		chkr = fmt.Sprintf("%s%c", chkr, c)

		if _, in := wordDict[chkr]; in {
			chkr = ""
			currStart = index + 1
		}

	}
	return currStart == len(s)
}

func main() {
	fmt.Printf("can break word %v\n", wordBreak("leetcode", []string{"leet", "code"}))
}
