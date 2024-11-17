package trie

import "strings"

//	type Replacer struct {
//		replacer *strings.Replacer
//	}

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

var StandardReplacer = strings.NewReplacer(standard...)
var TechnicalReplacer = strings.NewReplacer(technical...)

// var replacer = strings.NewReplacer(StandardReplacer...)
