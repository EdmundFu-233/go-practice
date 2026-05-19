package datastructures

// TrieNode represents a node in the Trie (prefix tree).
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// NewTrieNode creates a new TrieNode.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// Trie implements a prefix tree for efficient string search operations.
type Trie struct {
	root *TrieNode
	size int
}

// NewTrie creates a new empty Trie.
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
		size: 0,
	}
}

// Insert adds a word to the trie.
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	if !node.isEnd {
		node.isEnd = true
		t.size++
	}
}

// Search returns true if the word exists in the trie.
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith returns true if any word in the trie has the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// Delete removes a word from the trie. Returns true if the word was found and removed.
func (t *Trie) Delete(word string) bool {
	return t.deleteRecursive(t.root, word, 0)
}

func (t *Trie) deleteRecursive(node *TrieNode, word string, depth int) bool {
	if node == nil {
		return false
	}
	if depth == len(word) {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		t.size--
		return true
	}
	ch := rune(word[depth])
	child, exists := node.children[ch]
	if !exists {
		return false
	}
	deleted := t.deleteRecursive(child, word, depth+1)
	if deleted && len(child.children) == 0 && !child.isEnd {
		delete(node.children, ch)
	}
	return deleted
}

// Size returns the number of words in the trie.
func (t *Trie) Size() int {
	return t.size
}

// GetAllWithPrefix returns all words that start with the given prefix.
func (t *Trie) GetAllWithPrefix(prefix string) []string {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return nil
		}
		node = node.children[ch]
	}
	var results []string
	t.collect(node, prefix, &results)
	return results
}

func (t *Trie) collect(node *TrieNode, prefix string, results *[]string) {
	if node.isEnd {
		*results = append(*results, prefix)
	}
	for ch, child := range node.children {
		t.collect(child, prefix+string(ch), results)
	}
}
