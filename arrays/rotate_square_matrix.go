package arrays

import "fmt"

/*
Given an image represented by an NxN matrix,
where each pixel in the image is 4 bytes,
write a method to rotate the image by 90 degrees.
Can you do this in place?

I like to show this in a picture so that I can spot
the pattern.

For example:

This can represent the 2x2 image array.
[

	[ a, b ],
	[ c, d ]

]

What we want is:
[

	[ c, a ],
	[ d, b ]

]

The coordinates change from, to:

00 -> 01
01 -> 11
11 -> 10
10 -> 00
*/
type Pixel struct {
	r byte
	g byte
	b byte
	a byte
}

func rotateImage90Degress(image *[][]Pixel) {
	rows := len(*image)
	rotatedImage := make([][]Pixel, rows)
	// Make the destination NxN matrix that we'll
	// move the values into
	for row := 0; row < rows; row++ {
		if len((*image)[row]) != rows {
			panic(fmt.Sprintf("Non-square image with %d rows has a column with %d elements", rows, len((*image)[row])))
		}
		rotatedImage[row] = make([]Pixel, rows)
	}
	n := rows

	// Now that we have the destination matrix
	// let's traverse the src matrix to correctly
	// move the values into the src
	for x1 := 0; x1 < n; x1++ {
		y2 := x1
		for y1 := n - 1; y1 >= 0; y1-- {
			x2 := (n - 1) - y1
			rotatedImage[y2][x2] = (*image)[y1][x1]
		}
	}
	(*image) = rotatedImage
}

/*
If we have to rotate it in place, we have to determine where to swap each value.
The first thought that comes to mind is the linear algebra method of transposing
the matrix.

We would take the 'diagonal row' and 'reverse' it for each diagonal row in the
NxN matrix. Then, at the end, we can reverse each row, and the matrix is rotated.
*/
func rotateImage90DegressInPlace1(image *[][]Pixel) {
	rows := len(*image)
	// check for a non-square matrix
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

/*
Finally, we can pretend that, at each layer, there are 4 arrays that
we need to swap each element at, and then offset inwards to the
next level of arrays.

For example,

[
	[ a, b, c, d ],
	[ e, f, g, h ],
	[ i, j, k, l ],
	[ m, n, o, p ]
]

We have the arrays:

1 (top)
[ a, b, c ]
2 (right)
[ d,
  h,
  l
]
3. (bottom)
[ p, o, n ]
4. (left)
[ m,
  i,
  e
]

Get the coordinates of each array start, and then traverse the arrays.
Put the top value into a temp variable, and then put the left value
into the top, the bottom value into the left, the right value into
the bottom, and then the top value (stored in temp) into the right.

Update the index into each array by one. Once you've reached the end
of each array, reduce n by 2 and increment the offset so that you're
at the next inner level of the matrix.
*/

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
