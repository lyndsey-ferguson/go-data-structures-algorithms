package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPairs(t *testing.T) {
	text := "bluebirdskyscraper"

	words := []string{
		"blue", "bird", "sky",
	}

	result := indexPairs(text, words)
	assert.Equal(t, [][]int{{0, 3}, {4, 7}, {8, 10}}, result)

	text = "programmingisfun"
	words = []string{"pro", "is", "fun", "gram"}
	result = indexPairs(text, words)
	assert.Equal(t, [][]int{{0, 2}, {3, 6}, {11, 12}, {13, 15}}, result)

	text = "interstellar"
	words = []string{"stellar", "star", "inter"}
	result = indexPairs(text, words)
	assert.Equal(t, [][]int{{0, 4}, {5, 11}}, result)

	text = "applepie"
	words = []string{"apple", "pie", "app"}
	result = indexPairs(text, words)
	assert.Equal(t, [][]int{{0, 2}, {0, 4}, {5, 7}}, result)

	text = "datastructuresandalgorithms"
	words = []string{"data", "struct", "algo"}
	result = indexPairs(text, words)
	assert.Equal(t, [][]int{{0, 3}, {4, 9}, {17, 20}}, result)
}
