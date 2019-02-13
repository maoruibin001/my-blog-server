package main

import (
	"fmt"
	"reflect"
	"sort"
)

func getType(a interface{}) string {
	k := reflect.TypeOf(a)
	return fmt.Sprint(k)
}
func main() {
	results := []int{1, 3, 5}
	sort.Slice(results, func(i, j int) bool {

	fmt.Println(i, j)
		return false
	})
}