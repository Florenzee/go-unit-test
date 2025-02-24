package logic_2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal6(t *testing.T) {
	result := Soal6(9)

	//cek ukuran matriks
	assert.Equal(t, len(result), 9)    //matriks ukuran 9*9
	assert.Equal(t, result[8][8], 205) //elemen di baris ke-9 dan kolom ke-9 (index 8,8) adalah 18

	//cek value matriks
	resultVal := Soal6(9)
	expectedVal := [][]int{
		{1, 4, 7, 10, 13, 16, 19, 22, 25},
		{44, 42, 40, 38, 36, 34, 32, 30, 28},
		{46, 49, 52, 55, 58, 61, 64, 67, 70},
		{89, 87, 85, 83, 81, 79, 77, 75, 73},
		{91, 94, 97, 100, 103, 106, 109, 112, 115},
		{134, 132, 130, 128, 126, 124, 122, 120, 118},
		{136, 139, 142, 145, 148, 151, 154, 157, 160},
		{179, 177, 175, 173, 171, 169, 167, 165, 163},
		{181, 184, 187, 190, 193, 196, 199, 202, 205}}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
