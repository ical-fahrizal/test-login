package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	Logic1()
}

func Logic1() {
	log.Printf(`Input the number of families : `)
	reader := bufio.NewReader(os.Stdin)
	inputTotalFamily, _ := reader.ReadString('\n')
	inputTotalFamily = strings.TrimSpace(inputTotalFamily)
	numFamily, err := strconv.ParseInt(inputTotalFamily, 10, 64)
	if err != nil {
		log.Printf(`Not Data`)
		return
	}

	log.Println(`Input the number of members in the family`)
	log.Printf(`(separated by a space) : `)
	var arrCountFamily []int

	for i := 0; i <= int(numFamily)-1; i++ {
		countFamily, _ := reader.ReadString('\n')
		countFamily = strings.TrimSpace(countFamily)
		numCountFamily, err := strconv.ParseInt(countFamily, 10, 64)
		if err != nil {
			log.Printf(`Not Data`)
			continue
		}
		if numCountFamily > 4 {
			log.Printf(`One family can't have more than 4`)
			return
		}
		arrCountFamily = append(arrCountFamily, int(numCountFamily))
	}
	sort.Ints(arrCountFamily)

	jmlBus := 0
	for i := len(arrCountFamily) - 1; i > 0; i-- {
		if arrCountFamily[i] == 4 {
			arrCountFamily[i] = 0
			jmlBus += 1
		} else if arrCountFamily[i] == 0 {
			continue
		} else {
			for j := i - 1; j > 0; j-- {
				if arrCountFamily[j] == 0 {
					continue
				}
				jmlPenumpang := arrCountFamily[i] + arrCountFamily[j]
				if jmlPenumpang <= 4 {
					arrCountFamily[i] = 0
					arrCountFamily[j] = 0
					break
				}
			}
			jmlBus += 1
		}
	}

	log.Printf(`Minimum bus required is : %v`, jmlBus)
}
