package tries

import "fmt"

// IndexPairs finds all index pairs of substrings in the text that match any word in the trie
func indexPairs(text string, words []string) [][]int {
	var result [][]int

	t := &TrieNode{}
	for _, w := range words {
		t.Insert(w)
	}
	/*
		we have our trie

		- initialize an index to 0; initialize node to the root of the trie
		- get text[index] as a charIndex
		- if node.children[charIndex] != nil then
			- if rangeStart == -1; set rangeStart to index
			- setRangeEnd to index
			- set node to node.children[charIndex]
			- if node.IsEnd, then append to result && reset node = t && rangeStart = -1
	*/
	rangeEnd := -1

	for i := 0; i < len(text); i++ {
		charIndex := characterIndex(text[i])
		rangeStart := -1
		node := t.Children[charIndex]
		if node == nil {
			continue
		}
		for j := i + 1; j < len(text); j++ {
			charIndex := characterIndex(text[j])
			child := node.Children[charIndex]
			if child == nil {
				break // out of j loop
			}
			if rangeStart == -1 {
				rangeStart = i
			}
			rangeEnd = j
			node = child
			if node.IsEnd {
				fmt.Printf("found word '%v'\n", text[rangeStart:rangeEnd+1])
				result = append(result, []int{rangeStart, rangeEnd})
			}
		}
	}

	return result
}
