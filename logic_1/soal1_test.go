package logic_1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal1(t *testing.T) {
	//cek panjang index
	result := Soal1(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[3], 7)
	assert.NotEqual(t, result[3], 9)

	result = Soal1(13)
	assert.NotEqual(t, result[3], 9)

	//cek value setiap index
	resultVal := Soal1(10)
	expectedVal := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
