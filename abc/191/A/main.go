package main

import "fmt"

func main() {
	var v, t, s, d int
	fmt.Scan(&v, &t, &s, &d)

	if v*t <= d && d <= v*s {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
