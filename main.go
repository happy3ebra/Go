package main

import "fmt"

//type int
//func expression() bool {
//	x := rand.Intn(10)
//	return x != 5
//
//}

func main() {

	//var i int
	//rand.Seed(time.Now().Unix())
	x := [9]int{0, 167, 267, 5656, 5657, 555, 555, 6567567, 777}
	for _, v := range x {
		fmt.Println(v)

	}

	for i := 0; i < len(x); i += 2 {
		fmt.Println(x[i])
	}

}
