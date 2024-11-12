package trie

import (
	"fmt"
	"strings"
)

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

var newNode = NewTrie

func (current *Node) Add(word string) {
	for _, letter := range strings.ToLower(word) {
		if _, ok := current.children[letter]; !ok {
			current.children[letter] = newNode()
		}
		current = current.children[letter]
	}
	current.isEnd = true
}

func (current *Node) Search(word string) bool {
	for _, letter := range strings.ToLower(word) {
		if _, ok := current.children[letter]; !ok {
			return false
		}
		current = current.children[letter]
	}
	return current.isEnd
}

func (current *Node) StartsWith(word string) {

}
func (current *Node) EndsWith(word string) {

}
func (current *Node) Contains(word string) int {
	initial := rune(word[0])
	for range current.children {
		if _, ok := current.children[initial]; !ok {
			continue
		}
		for _, letter := range strings.ToLower(word) {
			_, ok := current.children[letter]
			if !ok {
				continue
			}
			current = current.children[letter]
		}
		fmt.Printf("Child count: %d, %s\n", len(current.children), word)
		return len(current.children)
	}
	return 0
}
