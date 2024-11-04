package tries

type TrieNode struct {
	Children [26]*TrieNode // Represents each letter of the alphabet.
	IsEnd    bool          // Flag to represent if the node is the end of a word.
}

func characterIndex(c byte) int {
	return int(c) - int('a')
}

// Inserts a word into the trie.
func (t *TrieNode) Insert(word string) {
	if t == nil {
		t = &TrieNode{}
	}
	node := t
	for i := 0; i < len(word); i++ {
		childIndex := characterIndex(word[i])
		if node.Children[childIndex] == nil {
			node.Children[childIndex] = &TrieNode{}
		}
		node = node.Children[childIndex]
	}
	node.IsEnd = true
}

// Returns if the word is in the trie.
func (t *TrieNode) Search(word string) bool {
	node := t

	for i := 0; i < len(word); i++ {
		childIndex := characterIndex(word[i])
		if node.Children[childIndex] == nil {
			return false
		}
		node = node.Children[childIndex]
	}
	return node.IsEnd
}

// Returns if there is any word in the trie that starts with the given prefix.
func (t *TrieNode) StartsWith(prefix string) bool {
	node := t

	for i := 0; i < len(prefix); i++ {
		childIndex := characterIndex(prefix[i])
		if node.Children[childIndex] == nil {
			return false
		}
		node = node.Children[childIndex]
	}
	return true
}
