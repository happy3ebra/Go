package main

import (
	"fmt"
	"math/rand"
)

//type int
func expression() bool {
	x := rand.Intn(10)
	return x != 5

}

func main() {

	//var a int
	for expression() {
		//a := 0
		fmt.Println("Hello World")

	}
}
