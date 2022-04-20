package main

import "fmt"

func HelloWorld(name string, s string) {
	fmt.Println("Hello WOrld", name, s)
}

func main() {
	name := "Aleksei"
	s := "Trivium"
	HelloWorld(name, s)
}
