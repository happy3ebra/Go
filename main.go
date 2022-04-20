package main

import (
	"fmt"
)

//type int
//func expression() bool {
//	x := rand.Intn(10)
//	return x != 5
//
//}

func main() {

	//var i int
	//rand.Seed(time.Now().Unix())
	var i int

	for i := 0; i < 5; i++ {

		if i == 3 {
			continue
		}
		//a := 0
		fmt.Println("Hello World", i)

	}
	fmt.Println(i)
}
