package main

import (
	"fmt"

	"tudai2021.com/model"
)

func main() {

	str := "TX04ABCD"
	r, err := model.GenerarRdo(str)

	if err == nil {
		fmt.Println(r)
	}
}
