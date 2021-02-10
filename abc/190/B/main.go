package main

import "fmt"

func main() {
	var n, s, d int
	fmt.Scan(&n, &s, &d)

	x := make([]int, n)
	y := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	for i := 0; i < n; i++ {
		if x[i] < s && d < y[i] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
