package logic_2

func Soal1(n int) (result [][]int) {
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)

		//mengisi setiap matriks dengan nilai sesuai soal
		for j := 0; j < n; j++ {
			result[i][j] = 1 + 2*j
		}
	}
	return result
}
