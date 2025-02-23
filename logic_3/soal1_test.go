package logic_3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal1(t *testing.T) {
	result := Soal1(9)

	//cek ukuran matriks
	assert.Equal(t, len(result), 9)   //matriks ukuran 9*9
	assert.Equal(t, result[8][8], 89) //elemen di baris ke-9 dan kolom ke-9 (index 8,8) adalah 18

	//cek value matriks
	resultVal := Soal1(9)
	expectedVal := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{5, 3, 0, 0, 0, 0, 0, 0, 0},
		{7, 9, 11, 0, 0, 0, 0, 0, 0},
		{19, 17, 15, 13, 0, 0, 0, 0, 0},
		{21, 23, 25, 27, 29, 0, 0, 0, 0},
		{41, 39, 37, 35, 33, 31, 0, 0, 0},
		{43, 45, 47, 49, 51, 53, 55, 0, 0},
		{71, 69, 67, 65, 63, 61, 59, 57, 0},
		{73, 75, 77, 79, 81, 83, 85, 87, 89}}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
