package model

import (
	"strconv"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(t string, l int, v string) Result {
	return Result{t, l, v}
}

func validateChar(str string) bool {

	for _, ch := range str {
		if !(ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z') {
			return false
		}
	}
	return true
}

func getType(str string) string { // Devuelve los primeros 2 caracteres de la cadena como t
	return str[:2]
}

func getLength(str string) string { // Devuelve la cantidad de caracteres pasado a enteros
	return str[2:4]

}
func getValue(str string) string { // Devuelve los caracteres a partir de la posicion 3 de la cadena
	return str[4:]
}

func GenerarRdo(s string) Result {
	var r Result
	if len(s) >= 4 {
		t := getType(s)
		if validateChar(t) {
			l := getLength(s)
			sv, err := strconv.Atoi(l)
			if err == nil {
				v := getValue(s)
				if validateChar(v) {
					count := len(v)
					if sv == count {
						r.Type = t
						r.Length = sv
						r.Value = v
					}
				}
			}
		}
	}
	return r
}
