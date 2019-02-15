package main

import (
	"fmt"
	"reflect"
)

func getType(a interface{}) string {
	k := reflect.TypeOf(a)
	return fmt.Sprint(k)
}

func getA() interface{} {
	return []string{"1"}
}
func main() {
	var a []string

	fmt.Println(len(a))
}