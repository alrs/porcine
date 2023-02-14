package main

import (
	"bufio"
	"log"
	"os"

	"github.com/alrs/phonetic"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		_, err := os.Stdout.Write(phonetic.BytesToNATO(scanner.Bytes()))
		if err != nil {
			log.Fatal(err)
		}
	}
}
