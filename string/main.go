package main

import (
	"fmt"
	"strings"
)

func main() {

	//字符串转数组切边
	s := "a,b,c,d"
	sArray := strings.Split(s, ",")
	fmt.Println(sArray)

	//数组切边转数组
	a := []string{"asdfa", "asdfasfd"}
	fmt.Println(strings.Join(a, ","))

}
