package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	t := make(map[int]bool)
	for i := 0; i < n; i++ {
		var d int
		fmt.Scan(&d)
		t[d] = true
	}
	fmt.Println(len(t))
}
