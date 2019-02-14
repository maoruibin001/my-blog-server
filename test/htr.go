package main

import (
	"fmt"
	"reflect"
	"time"
)

func getType(a interface{}) string {
	k := reflect.TypeOf(a)
	return fmt.Sprint(k)
}

func getA() interface{} {
	return []string{"1"}
}
func main() {
	var a = time.Now().Format("2006年01月02日 15时04分05秒")
	fmt.Println(a)
}