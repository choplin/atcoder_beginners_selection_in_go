package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var prevT, prevX, prevY int
	for i := 0; i < n; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)

		diffT, diffX, diffY := t-prevT, x-prevX, y-prevY
		if diffX < 0 {
			diffX = -diffX
		}
		if diffY < 0 {
			diffY = -diffY
		}
		if diffT < diffX+diffY || (diffT+diffX+diffY)%2 != 0 {
			fmt.Println("No")
			return
		}
		prevT, prevX, prevY = t, x, y
	}
	fmt.Println("Yes")
}
