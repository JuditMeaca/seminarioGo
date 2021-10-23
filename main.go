package main

import (
	"fmt"

	"tudai2021.com/model"
)

func main() {

	str := "TX04ABCD"
	r := model.GenerarRdo(str)

	if r.Length == 0 && (len(r.Value)) == 0 {
		fmt.Println("Error de cadena")
	} else {
		fmt.Println(r)
	}
}
