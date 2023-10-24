package arrays

import "fmt"

type Pixel struct {
	r byte
	g byte
	b byte
	a byte
}

func rotateImage90Degress(image *[][]Pixel) {
	rows := len(*image)
	rotatedImage := make([][]Pixel, rows)
	for row := 0; row < rows; row++ {
		if len((*image)[row]) != rows {
			panic(fmt.Sprintf("Non-square image with %d rows has a column with %d elements", rows, len((*image)[row])))
		}
		rotatedImage[row] = make([]Pixel, rows)
	}
	n := rows

	for x1 := 0; x1 < n; x1++ {
		y2 := x1
		for y1 := n - 1; y1 >= 0; y1-- {
			x2 := (n - 1) - y1
			rotatedImage[y2][x2] = (*image)[y1][x1]
		}
	}
	(*image) = rotatedImage
}

func rotateImage90DegressInPlace1(image *[][]Pixel) {
	rows := len(*image)
	for row := 0; row < rows; row++ {
		if len((*image)[row]) != rows {
			panic(fmt.Sprintf("Non-square image with %d rows has a column with %d elements", rows, len((*image)[row])))
		}
	}
	n := rows

	// transpose matrix
	for x := 0; x < n; x++ {
		for y := x; y < n; y++ {
			(*image)[y][x], (*image)[x][y] = (*image)[x][y], (*image)[y][x]
		}
	}
	// reverse each row
	for x := 0; x < n/2; x++ {
		for y := 0; y < n; y++ {
			(*image)[y][x], (*image)[y][n-1-x] = (*image)[y][n-1-x], (*image)[y][x]
		}
	}
}

func rotateImage90DegressInPlace2(image *[][]Pixel) {

	offset := 0
	for n := len(*image); n > 1; n = n - 2 {
		for i := 0; i < n-1; i++ {
			topY, topX := offset, offset+i
			leftY, leftX := offset+n-1-i, offset
			bottomY, bottomX := offset+n-1, offset+n-1-i
			rightY, rightX := offset+i, offset+n-1

			tmp := (*image)[topY][topX]
			(*image)[topY][topX] = (*image)[leftY][leftX]
			(*image)[leftY][leftX] = (*image)[bottomY][bottomX]
			(*image)[bottomY][bottomX] = (*image)[rightY][rightX]
			(*image)[rightY][rightX] = tmp
		}
		offset += 1
	}

}

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
