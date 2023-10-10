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
