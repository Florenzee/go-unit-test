package logic_1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal7(t *testing.T) {
	//cek panjang index
	resultIndex1 := Soal7(10)
	assert.Equal(t, len(resultIndex1), 10)

	resultIndex2 := Soal7(11)
	assert.Equal(t, len(resultIndex2), 11)

	//cek value setiap index
	resultVal1 := Soal7(10)
	expectedVal1 := []int{1, 3, 5, 7, 9, 9, 7, 5, 3, 1}
	assert.Equal(t, expectedVal1, resultVal1)
	fmt.Printf("result: %v\n", resultVal1)

	resultVal2 := Soal7(11)
	expectedVal2 := []int{1, 3, 5, 7, 9, 11, 9, 7, 5, 3, 1}
	assert.Equal(t, expectedVal2, resultVal2)
	fmt.Printf("result: %v\n", resultVal2)
}
