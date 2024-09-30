package main

import "fmt"

func main() {
	var a, b, n, ans int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	for i := 1; i <= n; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		if a <= sum && sum <= b {
			ans += i
		}
	}
	fmt.Println(ans)
}
