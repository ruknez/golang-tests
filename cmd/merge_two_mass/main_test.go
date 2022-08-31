package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeMass(t *testing.T) {
	inputData := []struct {
		a []int
		b []int
	}{
		{
			a: []int{0, 2, 4, 7},
			b: []int{3, 4, 7, 15, 32},
		},
		{
			a: nil,
			b: []int{3, 4, 7, 15, 32},
		},
		{
			a: []int{0, 2, 4, 7},
			b: []int{0, 2, 4, 7},
		},
	}

	for _, data := range inputData {
		myRes := mergeMass(data.a, data.b)
		totalRes := append(data.a, data.b...)
		sort.Ints(totalRes)
		assert.Equal(t, totalRes, myRes)
	}
}
