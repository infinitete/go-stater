package main

import . "github.com/infinitete/go-stater/common"
import "fmt"

func main() {
	var person = MakePerson("renshan", 29)

	fmt.Println(person.GetName())
}
