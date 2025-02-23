package logic_2

func Soal4(n int) (result [][]int) {
	result = make([][]int, n)

	value := 1

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)

		for j := 0; j < n; j++ {
			result[i][j] = value
			if j == 0 || i == 0 || j == n-1 {
				value += 3
			} else {
				value += 2
			}
		}
	}
	return result
}
