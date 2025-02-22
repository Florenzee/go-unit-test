package logic_2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal1(t *testing.T) {
	result := Soal1(10)

	//cek ukuran matriks
	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[9][9], 19)

	//cek value matriks
	resultVal := Soal1(9)
	expectedVal := [][]int{
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{1, 3, 5, 7, 9, 11, 13, 15, 17}}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
