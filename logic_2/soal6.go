package logic_2

func Soal6(n int) (result [][]int) {
	result = make([][]int, n)

	value := 1

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)

		for j := 0; j < n; j++ {

			if i%2 == 0 {
				result[i][j] = value
				value += 3
			} else {
				result[i][n-1-j] = value
				value += 2
			}
		}
	}
	return result
}
