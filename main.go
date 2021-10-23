package main

import (
	"fmt"

	"tudai2021.com/model"
)

func main() {
	str := "TX04ABC"
	r := model.GenerarRdo(str)

	fmt.Println(r)

}
