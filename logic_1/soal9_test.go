package logic_1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal9(t *testing.T) {
	//cek panjang index
	resultIndex1 := Soal9(10)
	assert.Equal(t, len(resultIndex1), 10)

	resultIndex2 := Soal9(11)
	assert.Equal(t, len(resultIndex2), 11)

	//cek value setiap index
	resultVal1 := Soal9(10)
	expectedVal1 := []int{3, 6, 9, 12, 15, 15, 12, 9, 6, 3}
	assert.Equal(t, expectedVal1, resultVal1)
	fmt.Printf("result: %v\n", resultVal1)

	resultVal2 := Soal9(11)
	expectedVal2 := []int{3, 6, 9, 12, 15, 18, 15, 12, 9, 6, 3}
	assert.Equal(t, expectedVal2, resultVal2)
	fmt.Printf("result: %v\n", resultVal2)
}
