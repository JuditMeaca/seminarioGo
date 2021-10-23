package model

import (
	"errors"
	"fmt"
	"strconv"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(t string, long int, val string) Result {
	return Result{t, long, val}
}

func getType(str string) string { // Devuelve los primeros 2 caracteres de la cadena como tipo
	return str[:2]
}

func getValue(str string) string { // Devuelve los caracteres a partir de la posicion 3 de la cadena
	return str[4:]
}

func getLength(str string) string { // Devuelve la cantidad de caracteres pasado a enteros

	return str[2:4]
}

func validateChar(s string) bool {

	for _, ch := range s {
		if !(ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z') {
			return false
		}
	}
	return true
}

func GenerateResult(s string) (Result, error) {
	var r Result

	if len(s) >= 4 { // Pregunta si el string es mayor o igual a 4 caracteres

		t := getType(s)      // Asigna a t lo que devuelve la funcion getType
		if validateChar(t) { // Valida que t contenga caracteres validos

			l := getLength(s)            //le asigna a l lo que retorna la funcionalidad getLength
			sv, error := strconv.Atoi(l) // asigna a sv el l convertido a int
			if error == nil {
				v := getValue(s)     //asigna a v lo que devuelve getValue
				if validateChar(v) { // lo valida

					count := len(v) //guarda en count la cantidad de char de v en numeros
					fmt.Println(t)
					if sv == count {
						r.Type = t

						r.Length = sv
						r.Value = v

					}

				}

			}
		}

	}
	return r, errors.New("cadena incorrecta")
}
