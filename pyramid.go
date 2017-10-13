package main

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// smallest obtains the smallest integer from an array
func smallest(n []int) int {
	sort.Ints(n)
	return n[0]
}

// shortest computes the smallest value that is
// obtained by sliding through the pyramid's levels.
// Thanks to:
//
//		Sam Hughes: https://github.com/srh
//
func shortest(pyramid [][]int) int {
	cur := []int{0}
	for i := range pyramid {
		row := make([]int, len(pyramid[i]))

		for j := range pyramid[i] {
			var shortp int

			switch {
			case j-1 < 0:
				shortp = cur[j]

			case j == len(cur):
				shortp = cur[j-1]

			default:
				shortp = min(cur[j], cur[j-1])
			}

			row[j] = shortp + pyramid[i][j]
		}

		cur = row
	}

	return smallest(cur)
}

func generate(levels int, min, max int) (ns [][]int) {
	ns = append(ns, []int{levels})

	for lv := 1; lv <= levels; lv++ {
		var nlv []int

		for i := 0; i < lv; i++ {
			nlv = append(nlv, randInt(min, max))
		}

		ns = append(ns, nlv)
	}

	return
}
