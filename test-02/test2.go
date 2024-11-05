package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// n := 5051
	log.Println(`Input Number: `)
	log.Println()
	reader := bufio.NewReader(os.Stdin)
	strInput, _ := reader.ReadString('\n')
	strInput = strings.TrimSpace(strInput)
	numInput, err := strconv.ParseInt(strInput, 10, 64)
	if err != nil {
		log.Printf(`Not Data`)
		return
	}

	// n = intInput
	hasil := reverseInteger(int(numInput))

	log.Printf(`%v`, hasil)
}

func reverseInteger(n int) int {

	reverse := 0

	for n > 0 {
		reverse = reverse*10 + n%10
		n = n / 10
	}

	return reverse
}
