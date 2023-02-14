// package phonetic
// Copyright (C) 2023 Lars Lehtonen KJ6CBE

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package phonetic

import (
	"bytes"
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
		'4': "foier",
		'5': "fife",
		'6': "six",
		'7': "seven",
		'8': "eight",
		'9': "nine",
		'a': "alpha",
		'b': "bravo",
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

func StringToNATO(s string) string {
	return stringToNATO(s)
}

func stringToNATO(s string) string {
	return string(bytesToNATO([]byte(s)))
}

func BytesToNATO(b []byte) []byte {
	return bytesToNATO(b)
}

func bytesToNATO(b []byte) []byte {
	result := make([]byte, 0, len(b))
	b = bytes.ToLower(b)
	for i, v := range b {
		if w, ok := byteNATO[v]; ok {
			result = append(result, w...)
			if i < len(b)-1 {
				result = append(result, []byte(" ")[0])
			}
		} else {
			result = append(result, v)
		}
	}
	return result
}
