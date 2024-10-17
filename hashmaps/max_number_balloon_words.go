package hashmaps

import (
	"math"
)

type Solution struct{}

func minInt(ints ...int) (min int) {
	min = math.MaxInt32
	for _, i := range ints {
		if i < min {
			min = i
		}
	}
	return
}

func maxNumberOfBalloons(text string) int {
	hash := make(map[rune]int, 7)
	hash['b'] = 0
	hash['a'] = 0
	hash['l'] = 0
	hash['o'] = 0
	hash['n'] = 0

	for _, r := range text {
		hash[r]++
	}

	min := minInt(
		hash['b'],
		hash['a'],
		hash['l']/2,
		hash['o']/2,
		hash['n'],
	)

	return min
}
