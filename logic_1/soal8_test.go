package logic_1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal8(t *testing.T) {
	//cek panjang index
	resultIndex1 := Soal8(10)
	assert.Equal(t, len(resultIndex1), 10)

	resultIndex2 := Soal8(11)
	assert.Equal(t, len(resultIndex2), 11)

	//cek value setiap index
	resultVal1 := Soal8(10)
	expectedVal1 := []int{2, 4, 6, 8, 10, 10, 8, 6, 4, 2}
	assert.Equal(t, expectedVal1, resultVal1)
	fmt.Printf("result: %v\n", resultVal1)

	resultVal2 := Soal8(11)
	expectedVal2 := []int{2, 4, 6, 8, 10, 12, 10, 8, 6, 4, 2}
	assert.Equal(t, expectedVal2, resultVal2)
	fmt.Printf("result: %v\n", resultVal2)
}
