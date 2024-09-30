package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int

	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)

	var ans int
	for i, v := range a {
		if i%2 == 0 {
			ans += v
		} else {
			ans -= v
		}
	}
	if len(a)%2 == 0 {
		ans *= -1
	}
	fmt.Println(ans)
}
