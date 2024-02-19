package main

import (
	"fmt"
	"time"
)

var num int = 0

func increment(ch chan bool) {

	for {
		time.Sleep(time.Second * 1)
		num += 1
		if num == 10 {
			ch <- true
		}
	}
}

func main() {
	ch := make(chan bool)

	go increment(ch)
	go increment(ch)

	<-ch

	fmt.Println("Number", num)
}
