package trie

import "strings"

var replacers = []string{
	".", " ",
	",", " ",
	"-", " ",
	"+", " ",
	"?", "",
	"!", "",
	":", " ",
	";", " ",
	"\n", " ",
	"\r", " ",
	"\t", " ",
	"'s", "",
	"[", "",
	"]", "",
	"(", "",
	")", "",
	"'", "",
	"‘", "",
	"\"", "",
	"'ll", " will",
}

var replacer = strings.NewReplacer(replacers...)
