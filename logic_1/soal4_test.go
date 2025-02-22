package logic_1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal4(t *testing.T) {
	//cek panjang index
	resultIndex := Soal4(10)
	assert.Equal(t, len(resultIndex), 10)

	//cek value setiap index
	resultVal := Soal4(10)
	expectedVal := []int{19, 17, 15, 13, 11, 9, 7, 5, 3, 1}
	assert.Equal(t, expectedVal, resultVal)
	fmt.Printf("result: %v\n", resultVal)
}
