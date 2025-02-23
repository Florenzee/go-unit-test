package logic_1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal6(t *testing.T) {
	//cek panjang index
	resultIndex := Soal6(10)
	assert.Equal(t, len(resultIndex), 10)

	//cek value setiap index
	resultVal := Soal6(10)
	expectedVal := []int{30, 27, 24, 21, 18, 15, 12, 9, 6, 3}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
