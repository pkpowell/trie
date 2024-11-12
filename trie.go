package trie

import (
	"fmt"
	"strings"
	"sync"
)

// type letter rune
type Node struct {
	children map[rune]*Node
	isEnd    bool
	mtx      *sync.RWMutex
}

func NewTrie() *Node {
	return &Node{
		children: make(map[rune]*Node),
		isEnd:    false,
		mtx:      new(sync.RWMutex),
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

// search for exact word
func (current *Node) Search(word string) bool {
	for _, letter := range strings.ToLower(word) {
		if _, ok := current.children[letter]; !ok {
			return false
		}
		current = current.children[letter]
	}
	return current.isEnd
}

// func (current *Node) StartsWith(word string) {

// }
// func (current *Node) EndsWith(word string) {

// }
// search for partial string. Returns number of matches
func (current *Node) Contains(word string) (matches int) {
	word = strings.ToLower(word)
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
			matches++
			current = current.children[letter]
		}

		fmt.Printf("Child count: %d, %s\n", matches, word)

	}
	return matches
}
