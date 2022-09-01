package main

import (
	"errors"
	"fmt"
)

var notFoundElement = errors.New("not found element")

func main() {
	res, err := binarySearch([]int{-100, -1, 1, 2, 3, 4, 5}, -10)
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("res = ", res)
	}
}

func binarySearch(m []int, element int) (int, error) {
	if len(m) == 0 {
		return 0, notFoundElement
	}
	left := 0
	right := len(m) - 1

	for left <= right {
		mid := (left + right) / 2

		if m[mid] == element {
			return mid, nil
		}
		if m[mid] < element {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0, notFoundElement
}
