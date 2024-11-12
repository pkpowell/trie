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
func (root *Node) Contains(word string) (matches int) {
	word = strings.ToLower(word)
	var scan func(*Node)
	// wordLen := len(word)
	scan = func(root *Node) {
		count := 0
		for r, node := range root.children {
			// fmt.Printf("Child node: %s\n", word)
			if byte(r) == word[count] {
				count++
				if count == len(word) {
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
		fmt.Printf("Child count: %d\n", count)
		// return matches
	}
	scan(root)
	return matches
}
