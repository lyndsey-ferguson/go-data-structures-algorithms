package arrays

/*
Write an algorithm such that if an element in an MxN matrix is 0,
its entire row and column are set to 0.

scan through all the rows, for each row, scan through
all columns, if zero found, add row to rows to zero and
add columns to columns to zero.

However, this is expensive, basically worst case O(n^2•m^2).
We have unecessary work, repeat looping back and down. Duplicate
work, sometimes I have already set an element to zero.

The best that I could hope for is O(n • m)
*/

// zero the column from the beginning to the current row
// we'll take care of zeroing out the rest of the column
// in our traversal of the matrix
func zeroBackColumn(matrix [][]int, row int, column int) {
	for r := 0; r < row; r++ {
		matrix[r][column] = 0
	}
}

// zero the row from the beginning to the current column
// we'll take care of zeroing out the rest of the row
// in our traversal of the matrix
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
				// we've discovered a zero in the row
				// let's clear the elements up to this one
				// an unfortunate cost, but we didnt' know about it
				// and then each time we move forward, knowing that
				// we are zeroing out the row, we just set the
				// next row element to 0 (see below)
				if !zeroRow {
					zeroBackRow(matrix, row, column)
					zeroRow = true
				}
				if !zeroColumns[column] {
					// if we haven't alread declared that this
					// column has to be zeroed out, let's
					// set the previous values in the column
					// to zero and then record that, from now
					// on, this column is to be set to 0
					zeroBackColumn(matrix, row, column)
					zeroColumns[column] = true
				}
			} else if zeroRow || zeroColumns[column] {
				// we've discovered that either we're zeroing out
				// the entire row, or this column has to be
				// zeroed out
				matrix[row][column] = 0
			}
		}
	}
}
