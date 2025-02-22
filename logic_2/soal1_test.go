package logic_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal1(t *testing.T) {
	result := Soal1(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[9][9], 19)
}
