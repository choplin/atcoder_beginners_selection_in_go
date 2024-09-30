package main

import "fmt"

func main() {
	var n, y int
	var a, b, c int

	fmt.Scanf("%d %d", &n, &y)

OUTER:
	for i := 0; i <= n; i++ {
		if i*10000 > y {
			break
		}
		for j := 0; j <= n-i; j++ {
			if i*10000+j*5000 > y {
				break
			}
			k := n - i - j
			if i*10000+j*5000+k*1000 == y {
				a = i
				b = j
				c = k
				break OUTER
			}
		}
	}
	if a+b+c == 0 {
		a = -1
		b = -1
		c = -1
	}
	fmt.Printf("%d %d %d\n", a, b, c)
}
