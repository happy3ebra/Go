package main

import "fmt"

func HelloWorld(prefix string, name ...string) {
	fmt.Println("Hello WOrld", name)
}

func main() {
	//name := "Aleksei"
	//s := "Trivium"
	HelloWorld("1")
}
