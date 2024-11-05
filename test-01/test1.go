package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.Println(`Input Kata: `)
	log.Println()
	reader := bufio.NewReader(os.Stdin)
	intInput, _ := reader.ReadString('\n')

	t := intInput

	log.Println(`Input Kata/karakter yang dicari : `)
	log.Println()
	intInput, _ = reader.ReadString('\n')

	s := intInput

	hasil := Is_subsequence(t, s)

	log.Printf(`output : %v`, hasil)
}

func Is_subsequence(input string, data string) bool {

	i := 0
	j := 0

	for i < len(input) && j < len(data) {

		if input[i] == data[j] {
			i += 1
		}
		j += 1
	}

	if i == len(input) {
		return true
	} else {
		return false
	}
}
