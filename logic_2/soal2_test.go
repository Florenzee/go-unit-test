package logic_2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal2(t *testing.T) {
	result := Soal2(9)

	//cek ukuran matriks
	assert.Equal(t, len(result), 9)   //matriks ukuran 9*9
	assert.Equal(t, result[8][8], 18) //elemen di baris ke-9 dan kolom ke-9 (index 8,8) adalah 18

	//cek value matriks
	resultVal := Soal2(9)
	expectedVal := [][]int{
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18},
		{2, 4, 6, 8, 10, 12, 14, 16, 18}}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
