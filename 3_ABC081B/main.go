package main

import "fmt"

func main() {
	var n int
	var a []int
	var cnt int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		a = append(a, tmp)
	}

OUTER:
	for {
		for i := 0; i < n; i++ {
			if a[i]%2 != 0 {
				break OUTER
			} else {
				a[i] = a[i] / 2
			}
		}
		cnt++
	}

	fmt.Println(cnt)
}
