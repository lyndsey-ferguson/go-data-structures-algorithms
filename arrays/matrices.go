package arrays

func zeroBackColumn(matrix [][]int, row int, column int) {
	for r := 0; r < row; r++ {
		matrix[r][column] = 0
	}
}

func zeroBackRow(matrix [][]int, row int, column int) {
	for c := 0; c < column; c++ {
		matrix[row][c] = 0
	}
}

func zeroOutAxisInMatrix(matrix [][]int) {
	zeroColumns := make([]bool, len(matrix[0]))

	for row := 0; row < len(matrix); row++ {
		zeroRow := false
		for column := 0; column < len(matrix[0]); column++ {
			if matrix[row][column] == 0 {
				if !zeroRow {
					zeroBackRow(matrix, row, column)
					zeroRow = true
				}
				if !zeroColumns[column] {
					zeroBackColumn(matrix, row, column)
					zeroColumns[column] = true
				}
			} else if zeroRow || zeroColumns[column] {
				matrix[row][column] = 0
			}
		}
	}
}
