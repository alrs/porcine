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

package main

import (
	"bufio"
	"log"
	"os"

	"github.com/porcine/alrs/phonetic"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		_, err := os.Stdout.Write(phonetic.NATO.ConvertBytes(scanner.Bytes(), false))
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write([]byte("\n"))
	}
}
