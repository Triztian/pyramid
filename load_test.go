package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

var challenges []string

func loadChallenge(b *testing.B, path string) string {
	b.Helper()

	d, err := ioutil.ReadFile(path)
	if err != nil {
		b.Fatal(err)
	}

	return string(d)
}

func BenchmarkParse(b *testing.B) {
	s := loadChallenge(b, "./testdata/challenge_3.txt")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := parse(s)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkScan(b *testing.B) {
	// s := loadChallenge(b, "./testdata/challenge_3.txt")
	s := loadChallenge(b, "./testdata/challenge_3.txt")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r := strings.NewReader(s)
		_, err := scanReader(r)
		if err != nil {
			b.Fatal(err)
		}
	}
}
