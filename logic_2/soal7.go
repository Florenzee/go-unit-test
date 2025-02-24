package logic_2

func Soal7(n int) (result [][]int) {
	result = make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)

		for j := 0; j < n; j++ {
			if i == j {
				result[i][j] = 2*j + 1
			}
		}
	}
	return result
}
