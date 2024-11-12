package trie

import "strings"

// type letter rune
type Node struct {
	children map[rune]*Node
	isEnd    bool
}

func NewTrie() *Node {
	return &Node{
		children: make(map[rune]*Node),
		isEnd:    false,
	}
}

func (root *Node) Add(word string) {
	current := root
	for _, letter := range strings.ToLower(word) {
		if _, ok := current.children[letter]; !ok {
			current.children[letter] = NewTrie()
		}
		current = current.children[letter]
	}
	current.isEnd = true
}

func (root *Node) Search(word string) bool {
	current := root
	for _, letter := range strings.ToLower(word) {
		if _, ok := current.children[letter]; !ok {
			return false
		}
		current = current.children[letter]
	}
	return current.isEnd
}
