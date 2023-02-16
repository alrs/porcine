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

type letterMap map[rune]string

// An Alphabet reprents a mapping of Alphabetic letters to their
// documented word-length representations as well as modified
// words optimized for speech synthesis.
type Alphabet struct {
	literal letterMap
	spoken  letterMap
}

// World War II Combined Communications Board phonetic alphabet.
var CCB Alphabet

// 1956 ICAO phonetic alphabet.
var NATO Alphabet

func init() {
	initNATO()
	initCCB()
}

func newAlphabet(l letterMap, s letterMap) Alphabet {
	a := Alphabet{}
	a.literal = l
	a.spoken = s
	return a
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
			if w, ok := a.literal[rune(v)]; ok {
				result = append(result, w...)
			}
		}
		if i < len(b)-1 {
			result = append(result, []byte(" ")[0])
		}
	}
	return result
}
