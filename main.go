package main

import "fmt"

func HelloWorld(name string) {
	fmt.Println("Hello WOrld", name)
}

func main() {
	name := "Aleksei"
	HelloWorld(name)
}
