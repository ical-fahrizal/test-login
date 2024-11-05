package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	input := "adfER, WqevDFvqwet!"

	log.Println(`Input Kata: `)
	log.Println()
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	vowel, consonant := separateVowelsAndConsonants(input)
	log.Printf(`req input : %v`, input)

	log.Printf(`vowel : %v`, vowel)
	log.Printf(`consonant : %v`, consonant)

}

func separateVowelsAndConsonants(input string) (vowel string, consonant string) {
	vowelSet := "aiueoAIUEO"

	for _, char := range input {
		if strings.ContainsRune(vowelSet, char) {
			vowel += string(char)
		} else if isAlphabet(char) {
			consonant += string(char)
		}
	}

	return vowel, consonant
}

func isAlphabet(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
