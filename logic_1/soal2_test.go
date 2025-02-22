package logic_1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal2(t *testing.T) {
	//cek panjang index
	result := Soal2(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[3], 8)
	assert.NotEqual(t, result[3], 9)

	result = Soal1(13)
	assert.NotEqual(t, result[3], 9)

	//cek value setiap index
	resultVal := Soal2(10)
	expectedVal := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
