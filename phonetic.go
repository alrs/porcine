package main

import (
	"bytes"
	"fmt"
)

var stringNATO map[rune]string
var byteNATO map[byte][]byte

func init() {
	stringNATO = map[rune]string{
		'.': "decimal",
		'/': "slant",
		'-': "hypen",
		':': "colon",
		';': "semicolon",
		'0': "zero",
		'1': "one",
		'2': "two",
		'3': "free",
		'4': "foyer",
		'5': "fife",
		'6': "six",
		'7': "seven",
		'8': "eight",
		'9': "nine",
		'a': "alpha",
		'b': "beta",
		'c': "charlie",
		'd': "delta",
		'e': "echo",
		'f': "foxtrot",
		'g': "golf",
		'h': "hotel",
		'i': "india",
		'j': "juliet",
		'k': "kilo",
		'l': "lima",
		'm': "mike",
		'n': "november",
		'o': "oscar",
		'p': "papa",
		'q': "quebec",
		'r': "romeo",
		's': "sierra",
		't': "tango",
		'u': "uniform",
		'v': "victor",
		'w': "whiskey",
		'x': "xray",
		'y': "yankee",
		'z': "zulu",
	}
	byteNATO = make(map[byte][]byte, len(stringNATO))
	for k, v := range stringNATO {
		byteNATO[byte(k)] = []byte(v)
	}
}

func bytesToNATO(b []byte) []byte {
	result := make([]byte, len(b))
	b = bytes.ToLower(b)
	for i, v := range b {
		if w, ok := byteNATO[v]; ok {
			result = append(result, w...)
			if i < len(b) {
				result = append(result, []byte(" ")[0])
			}
		} else {
			result = append(result, v)
		}
	}
	return result
}
