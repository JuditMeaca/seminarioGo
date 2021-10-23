package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"seminarioGo.com.com/model"
)

func TestGenerateResult(t *testing.T) {
	str := "AB01CD"
	result := model.NewResult("AB", 1, "CD")
	r := model.GenerateResult(str)
	assert.Equal(t, r, result, "Fallo al generar la instancia de la estructura")
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
		test := model.GenerateResult(testData.Input)
		assert.Equal(t, test.Type, testData.Type, testData.Success)
		assert.Equal(t, test.Value, testData.Value, testData.Success)
		assert.Equal(t, test.Length, testData.Length, testData.Success)
	}
}