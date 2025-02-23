package logic_3

func Soal1(n int) (result [][]int) {
	result = make([][]int, n)

	value := 1

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)

		if i%2 == 0 {
			for j := 0; j <= i; j++ {
				result[i][j] = value
				value += 2
			}
		} else {
			for j := i; j >= 0; j-- {
				result[i][j] = value
				value += 2
			}
		}
	}
	return result
}
