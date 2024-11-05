package main

import (
	"log"
)

func main() {

	arr := []int{1, 2, 2, 3, 4, 5, 6, 7, 4}

	total := isUniformLength(arr)

	log.Printf(`Jumlah data yg sama : %v`, total)
}

func isUniformLength(arr []int) int {

	iTotal := 0

	for len(arr) > 1 {
		dataCheck := arr[0]
		copyArr := removeIndex(arr, 0)
		arr = copyArr

		for j, v1 := range copyArr {
			if dataCheck == v1 {
				iTotal += 1
				arr = removeIndex(copyArr, j)
				break
			}
		}
	}

	return iTotal
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
