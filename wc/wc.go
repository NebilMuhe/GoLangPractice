package main

import (
	"fmt"
)

func main() {
	fmt.Println("Own Wc")

	var c1, c2, c3 string

	for true {
		fmt.Print(">> ")
		val, err := fmt.Scanln(&c1, &c2, &c3)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(val)

		fmt.Printf("%v %v %v\n", c1, c2, c3)
	}
}

// func main() {
// 	fmt.Println("input text:")
// 	var w1, w2, w3 string
// 	n, err := fmt.Scanln(&w1, &w2, &w3)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Printf("number of items read: %d\n", n)
// 	fmt.Printf("read line: %s %s %s-\n", w1, w2, w3)
// }
