package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchAfterInsertion(t *testing.T) {
	trie := &TrieNode{}

	trie.Insert("apple")
	assert.True(t, trie.Search("apple"))
	assert.True(t, trie.StartsWith("app"))
	trie.Insert("app")
	assert.True(t, trie.Search("apple"))

	trie = &TrieNode{}

	trie.Insert("apple")
	trie.Insert("banana")
	assert.True(t, trie.Search("apple"))
	assert.True(t, trie.StartsWith("ban"))

	trie = &TrieNode{}
	trie.Insert("pineapple")
	assert.True(t, trie.StartsWith("pine"))
	assert.False(t, trie.StartsWith("pineappleapple"))
}
