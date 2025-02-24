package logic_2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal11(t *testing.T) {
	result := Soal11(9)

	//cek ukuran matriks
	assert.Equal(t, len(result), 9)   //matriks ukuran 9*9
	assert.Equal(t, result[8][8], 17) //elemen di baris ke-9 dan kolom ke-9 (index 8,8) adalah 18

	//cek value matriks
	resultVal := Soal11(9)
	expectedVal := [][]int{
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{0, 3, 5, 7, 9, 11, 13, 15, 17},
		{0, 0, 5, 7, 9, 11, 13, 15, 17},
		{0, 0, 0, 7, 9, 11, 13, 15, 17},
		{0, 0, 0, 0, 9, 11, 13, 15, 17},
		{0, 0, 0, 0, 0, 11, 13, 15, 17},
		{0, 0, 0, 0, 0, 0, 13, 15, 17},
		{0, 0, 0, 0, 0, 0, 0, 15, 17},
		{0, 0, 0, 0, 0, 0, 0, 0, 17},
	}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
