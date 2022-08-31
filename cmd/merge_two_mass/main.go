package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 4, 7}
	b := []int{3, 4, 7, 15, 32}
	fmt.Println(mergeMass(a, b))

}

func mergeMass(a, b []int) []int {
	resultLen := len(a) + len(b)
	result := make([]int, 0, resultLen)
	countA := 0
	countB := 0

	for i := 0; i < resultLen; i++ {
		if countA >= len(a) {
			result = append(result, b[countB:]...)
			break
		}
		if countB >= len(b) {
			result = append(result, a[countA:]...)
			break
		}

		if a[countA] > b[countB] {
			result = append(result, b[countB])
			countB++
		} else {
			result = append(result, a[countA])
			countA++
		}
	}
	return result
}
