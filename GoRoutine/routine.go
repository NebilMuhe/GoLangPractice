package main

import (
	"fmt"
	"time"
)

var num int = 4

func increment(ch chan bool) {

	for {
		time.Sleep(time.Second * 1)
		num += 1
		if num == 20 {
			ch <- true
		}
	}
}

func main() {
	ch := make(chan bool)

	go increment(ch)
	go increment(ch)
	//channels

	<-ch

	fmt.Println("Number", num)
}
