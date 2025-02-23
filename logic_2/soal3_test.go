package logic_2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal3(t *testing.T) {
	result := Soal3(9)

	//cek ukuran matriks
	assert.Equal(t, len(result), 9)    //matriks ukuran 9*9
	assert.Equal(t, result[8][8], 161) //elemen di baris ke-9 dan kolom ke-9 (index 8,8) adalah 18

	//cek value matriks
	resultVal := Soal3(9)
	expectedVal := [][]int{
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{19, 21, 23, 25, 27, 29, 31, 33, 35},
		{37, 39, 41, 43, 45, 47, 49, 51, 53},
		{55, 57, 59, 61, 63, 65, 67, 69, 71},
		{73, 75, 77, 79, 81, 83, 85, 87, 89},
		{91, 93, 95, 97, 99, 101, 103, 105, 107},
		{109, 111, 113, 115, 117, 119, 121, 123, 125},
		{127, 129, 131, 133, 135, 137, 139, 141, 143},
		{145, 147, 149, 151, 153, 155, 157, 159, 161}}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
