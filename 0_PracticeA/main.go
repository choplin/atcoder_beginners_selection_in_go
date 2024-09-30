package main

import "fmt"

func main() {
	var a, b, c int
	var s string
	var err error

	_, err = fmt.Scanf("%d", &a)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Scanf("%d %d", &b, &c)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Scanf("%s", &s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d %s\n", a+b+c, s)
}
