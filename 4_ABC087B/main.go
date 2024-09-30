package main

import "fmt"

func main() {
	var a, b, c, x, cnt int
	fmt.Scanf("%d\n%d\n%d\n%d\n", &a, &b, &c, &x)

	for i := 0; i <= a; i++ {
		if 500*i > x {
			break
		}
		for j := 0; j <= b; j++ {
			if 500*i+100*j > x {
				break
			}
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k == x {
					cnt++
					break
				} else if 500*i+100*j+50*k > x {
					break
				}
			}
		}
	}
	fmt.Println(cnt)
}
