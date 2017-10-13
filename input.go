package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Default min and max values for generation.
const (
	DefaultMin int64 = 0
	DefaultMax int64 = 10000000
)

// parse slices and splits a string to
// obtain a valid Pyramid representation
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

// scanReader is a bufio.Scanner based implementation
// for parsing an input that is provided by an `io.Reader`
// object.
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

func writeTo(out io.Writer, ns [][]int) error {
	buf := bufio.NewWriter(out)
	buf.WriteString(fmt.Sprint(ns[0][0]))
	buf.WriteString("\n")
	for lv := 1; lv < len(ns); lv++ {
		for i := range ns[lv] {
			buf.WriteString(fmt.Sprint(ns[lv][i]))
			if i != len(ns[lv])-1 {
				buf.WriteString(" ")
			}
		}

		if lv != ns[0][0] {
			buf.WriteString("\n")
		}
	}

	return buf.Flush()
}
