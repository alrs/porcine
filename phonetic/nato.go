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

func initNATO() {
	NATO = newAlphabet(letterMap{
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
	}, letterMap{
		'4': "foewhur",
		'j': "juliet",
		'p': "pawpaw",
		'x': "ecksray",
	},
	)
}
