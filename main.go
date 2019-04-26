package main

import . "github.com/infinitete/go-stater/common"
import "fmt"

func main() {
	var person = Person{
		Name: "renshan",
		Age:  29,
	}

	fmt.Println(person)
}
