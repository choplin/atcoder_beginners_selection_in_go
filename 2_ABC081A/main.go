package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var count int
	for _, c := range s {
		if c == '1' {
			count++
		}
	}
	fmt.Println(count)
}
