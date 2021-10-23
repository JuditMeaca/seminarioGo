package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tudai2021.com/model"
)

func TestGenerarRdo(t *testing.T) {

	str := "AB04DEFG"
	rdo := model.NewResult("AB", 4, "DEFG")
	r := model.GenerarRdo(str)
	assert.Equal(t, r, rdo, "Error")

}

func TestParser(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {
		r := model.GenerarRdo(testData.Input)
		assert.Equal(t, r.Type, testData.Type, testData.Success)
		assert.Equal(t, r.Value, testData.Value, testData.Success)
		assert.Equal(t, r.Length, testData.Length, testData.Success)
	}
}
