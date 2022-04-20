package main

import (
	"fmt"
	"math/rand"
	"time"
)

//type int
func expression() bool {
	x := rand.Intn(10)
	return x != 5

}

func main() {

	//var a int
	rand.Seed(time.Now().Unix())
	for expression() {
		//a := 0
		fmt.Println("Hello World")

	}
}
