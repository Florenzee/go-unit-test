package logic_1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal5(t *testing.T) {
	//cek panjang index
	resultIndex := Soal5(10)
	assert.Equal(t, len(resultIndex), 10)

	//cek value setiap index
	resultVal := Soal5(10)
	expectedVal := []int{20, 18, 16, 14, 12, 10, 8, 6, 4, 2}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
