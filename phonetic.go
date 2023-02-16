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

// An Alphabet reprents a mapping of alphabetic letters to their
// documented word-length representations as well as modified
// words optimized for speech synthesis.
type Alphabet struct {
	letterMap map[rune]string
	spoken    map[rune]string
}

var NATO Alphabet

func init() {
	NATO.letterMap = map[rune]string{
		'.': "decimal",
		'/': "slant",
		'-': "hypen",
		':': "colon",
		';': "semicolon",
		'0': "zero",
		'1': "one",
		'2': "two",
		'3': "tree",
		'4': "fower",
		'5': "fife",
		'6': "six",
		'7': "seven",
		'8': "eight",
		'9': "niner",
		'a': "alfa",
		'b': "bravo",
		'c': "charlie",
		'd': "delta",
		'e': "echo",
		'f': "foxtrot",
		'g': "golf",
		'h': "hotel",
		'i': "india",
		'j': "juliett",
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

	// override letters poorly pronounced by speech synthesis
	NATO.spoken = map[rune]string{
		'4': "foewhur",
		'j': "juliet",
		'p': "pawpaw",
		'x': "ecksray",
	}
}

func (a *Alphabet) Convert(s string, spoken bool) string {
	b := a.ConvertBytes([]byte(s), spoken)
	return string(b)
}

// ConvertBytes changes a byteslice of characters into an expanded byteslice
// of their phonetic representation, either literal or modified for speech
// synthesis.
func (a *Alphabet) ConvertBytes(b []byte, spoken bool) []byte {
	result := make([]byte, 0, len(b))
	b = bytes.ToLower(b)
	for i, v := range b {
		// is the rune in the lettermap?
		s, ok := a.spoken[rune(v)]
		if spoken && ok {
			result = append(result, s...)
		} else {
			if w, ok := a.letterMap[rune(v)]; ok {
				result = append(result, w...)
			}
		}
		if i < len(b)-1 {
			result = append(result, []byte(" ")[0])
		}
	}
	return result
}
