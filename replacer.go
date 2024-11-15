package trie

import "strings"

//	type Replacer struct {
//		replacer *strings.Replacer
//	}
var StandardReplacer = strings.NewReplacer(standard...)
var TechnicalReplacer = strings.NewReplacer(technical...)
var standard = []string{
	".", " ",
	",", " ",
	"-", " ",
	"+", " ",
	"?", " ",
	"!", " ",
	":", " ",
	";", " ",
	"...", " ",
	// "\n", " ",
	// "\r", " ",
	"\t", " ",
	"'s", "",
	"[", " ",
	"]", " ",
	"(", " ",
	")", " ",
	"'", " ",
	"‘", " ",
	"’", " ",
	"\"", "",
	"'ll", "", //" will",
}
var technical = []string{
	// ".", " ",
	",", " ",
	// "-", " ",
	"+", " ",
	"?", " ",
	"!", " ",
	// ":", " ",
	";", " ",
	"...", " ",
	// "\n", " ",
	// "\r", " ",
	"\t", " ",
	"'s", "",
	// "[", " ",
	// "]", " ",
	"(", " ",
	")", " ",
	"'", " ",
	"‘", " ",
	"’", " ",
	"\"", "",
	// "'ll", "", //" will",
}

// var replacer = strings.NewReplacer(StandardReplacer...)
