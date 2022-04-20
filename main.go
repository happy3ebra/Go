package main

import (
	"errors"
	"fmt"
	"strings"
)

func HelloWorld(prefix string, name ...string) (string, error) {

	if prefix == "" {
		return "", errors.New("Префикс не должен быть пустым ")

	}

	return fmt.Sprintf("%s Hello W0rld %v ", prefix, strings.Join(name, ",")), nil
}

func main() {
	//name := "Aleksei"
	//s := "Trivium"
	fmt.Println("Введите ваше имя ")
	var name string
	n, err := fmt.Scanln(&name)
	if err != nil {

	}

	str, err := HelloWorld(">>>>", name)
	if err != nil {
		fmt.Println("Произошла ошибка", err.Error())
		return
	}
	fmt.Println(str)

}
