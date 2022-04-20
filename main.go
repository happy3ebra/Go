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
	x := [3]int{0, 1, 2}
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

}
