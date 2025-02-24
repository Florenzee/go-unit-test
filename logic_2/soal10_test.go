package logic_2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal10(t *testing.T) {
	result := Soal10(9)

	//cek ukuran matriks
	assert.Equal(t, len(result), 9)   //matriks ukuran 9*9
	assert.Equal(t, result[8][8], 17) //elemen di baris ke-9 dan kolom ke-9 (index 8,8) adalah 18

	//cek value matriks
	resultVal := Soal10(9)
	expectedVal := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 3, 0, 0, 0, 0, 0, 0, 0},
		{1, 3, 5, 0, 0, 0, 0, 0, 0},
		{1, 3, 5, 7, 0, 0, 0, 0, 0},
		{1, 3, 5, 7, 9, 0, 0, 0, 0},
		{1, 3, 5, 7, 9, 11, 0, 0, 0},
		{1, 3, 5, 7, 9, 11, 13, 0, 0},
		{1, 3, 5, 7, 9, 11, 13, 15, 0},
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
	}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
