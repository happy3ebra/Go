package main

import (
	"fmt"
	"strings"
)

func HelloWorld(prefix string, name ...string) (string, error) {
	return fmt.Sprintf("%s Hello W0rld %v ", prefix, strings.Join(name, ",")), nil
}

func main() {
	//name := "Aleksei"
	//s := "Trivium"

	str, err := HelloWorld("!", "a", "d", "e")
	if err != nil {
		fmt.Println("Произошла ошибка", err.Error())
		return
	}
	fmt.Println(str)

}
