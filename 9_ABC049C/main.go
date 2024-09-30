package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	words := []string{"dream", "dreamer", "erase", "eraser"}

	for {
		if s == "" {
			fmt.Println("YES")
			return
		}

		matched := false
		for _, word := range words {
			if len(s) < len(word) {
				continue
			}

			if s[len(s)-len(word):] == word {
				s = s[:len(s)-len(word)]
				matched = true
				break
			}
		}

		if !matched {
			fmt.Println("NO")
			return
		}
	}
}
