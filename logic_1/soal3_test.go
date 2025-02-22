package logic_1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal3(t *testing.T) {
	//cek panjang index
	resultIndex := Soal3(10)
	assert.Equal(t, len(resultIndex), 10)

	//cek value setiap index
	resultVal := Soal3(10)
	expectedVal := []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
