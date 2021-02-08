package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a < b || (c == 0 && a == b) {
		fmt.Println("Aoki")
	} else {
		fmt.Println("Takahashi")
	}
}
