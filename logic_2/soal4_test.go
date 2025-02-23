package logic_2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal4(t *testing.T) {
	result := Soal4(9)

	//cek ukuran matriks
	assert.Equal(t, len(result), 9)    //matriks ukuran 9*9
	assert.Equal(t, result[8][8], 185) //elemen di baris ke-9 dan kolom ke-9 (index 8,8) adalah 18

	//cek value matriks
	resultVal := Soal4(9)
	expectedVal := [][]int{
		{1, 4, 7, 10, 13, 16, 19, 22, 25},
		{28, 31, 33, 35, 37, 39, 41, 43, 45},
		{48, 51, 53, 55, 57, 59, 61, 63, 65},
		{68, 71, 73, 75, 77, 79, 81, 83, 85},
		{88, 91, 93, 95, 97, 99, 101, 103, 105},
		{108, 111, 113, 115, 117, 119, 121, 123, 125},
		{128, 131, 133, 135, 137, 139, 141, 143, 145},
		{148, 151, 153, 155, 157, 159, 161, 163, 165},
		{168, 171, 173, 175, 177, 179, 181, 183, 185}}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
