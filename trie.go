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
		_, ok := current.children[letter]
		if !ok {
			current.children[letter] = newNode()
		}
		current = current.children[letter]
	}
	current.isEnd = true
}

// search for exact word
func (current *Node) Search(word string) bool {
	for _, letter := range strings.ToLower(word) {
		_, ok := current.children[letter]
		if !ok {
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
func (root *Node) Contains(word string) (matches int) {
	letters := []rune(strings.ToLower(word))
	var scan func(*Node)
	wordLen := len(word)
	fmt.Printf("Searching for %s, %d\n", word, len(root.children))
	// r := []rune(word)
	count := 0
	scan = func(root *Node) {
		for r, node := range root.children {
			if r == letters[count] {
				// fmt.Printf("Child node: %c, %c\n", r, letters[count])
				count++
				if count == wordLen {
					matches++
					count = 0
				}
			} else {
				count = 0
			}
			scan(node)
			// for _, letter := range word {
			// 	n, ok := node.children[letter]
			// 	// fmt.Printf("Node %#v\n", node.children)
			// 	if !ok {
			// 		fmt.Printf("Letter %s not found in %s\n", string(letter), word)
			// 		scan(node)
			// 	}
			// 	count++
			// 	node = n
			// }
			// if count == len(word) {
			// 	matches++
			// }

		}
		if matches > 0 {
			fmt.Printf("word count: %d x %s\n", matches, word)
		}
		// fmt.Printf("word count: %d\n", matches)
		// return matches
	}
	scan(root)
	return matches
}
