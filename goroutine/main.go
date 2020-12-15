package main

import (
	"fmt"
)

type ListCh struct {
	Nums   int
	Slogan string
}

func say(s string, c chan ListCh) {

	for i := 0; i < 1000; i++ {
		goListCh := ListCh{
			Nums:   i,
			Slogan: s,
		}
		c <- goListCh
	}

	close(c)
}

func main() {

	ch := make(chan ListCh)
	go say("world", ch)
	for i := range ch {
		fmt.Printf("获取的信息：%+v\n", i)
	}

}
