package main

import (
	"fmt"
	"strings"
)

func HelloWorld(prefix string, name ...string) string {
	return fmt.Sprintf("%s Hello W0rld %s ", prefix, strings.Join(name, ","))
}

func main() {
	//name := "Aleksei"
	//s := "Trivium"
	fmt.Println(
		HelloWorld("1", "a", "d", "e"),
	)

}
