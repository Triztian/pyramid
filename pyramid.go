package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func smallest(n []int) int {
	sort.Ints(n)
	return n[0]
}

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

func parse(s string) ([][]int, error) {
	ns := [][]int{}
	var lv int
	for _, line := range strings.Split(s, "\n") {
		ns = append(ns, []int{})
		line = strings.Trim(line, " ")
		for _, item := range strings.Split(line, " ") {
			n, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				return nil, err
			}
			ns[lv] = append(ns[lv], int(n))
		}
		lv++
	}

	return ns, nil
}

func scanReader(r io.Reader) ([][]int, error) {
	var ns [][]int

	lsc := bufio.NewScanner(r)
	lineBuf := bytes.NewBuffer([]byte{})
	for lsc.Scan() {
		var lv []int

		lineBuf.Write(lsc.Bytes())
		nsc := bufio.NewScanner(lineBuf)
		nsc.Split(bufio.ScanWords)
		for nsc.Scan() {
			n, err := strconv.ParseInt(string(nsc.Bytes()), 10, 32)
			if err != nil {
				return nil, err
			}
			lv = append(lv, int(n))
		}
		lineBuf.Reset()

		ns = append(ns, lv)
	}

	return ns, nil
}

func main() {
	// https://www.reddit.com/r/dailyprogrammer/comments/6vi9ro/170823_challenge_328_intermediate_pyramid_sliding/?st=j8oziieq&sh=700700bb
	file := os.Args[1]

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	pyramid, err := scanReader(f)
	if err != nil {
		log.Fatalln(err)
	}

	ans := shortest(pyramid[1:])

	fmt.Println("Answer:", ans)
}
