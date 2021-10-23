package main

import (
	"fmt"

	"seminarioGo.com/model"
)

func main() {

	str := "TX03ABCD"

	r, error := model.GenerateResult(str)

	if error == nil {
		fmt.Println(r)
	} else {
		fmt.Println("Cadena incorrecta")
	}

}
