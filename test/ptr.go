package main

import "fmt"

func main() {
	var a = []int{1, 2, 4}

	b := &a[1]

	b = 20

	fmt.Println(b)

	fmt.Println(a)
}
