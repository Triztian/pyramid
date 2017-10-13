 # Benchmarks

 ## `scanReader` 
 ```
 λ pyramid go tool pprof -list='scanReader' cpu.out
Total: 1.11s
ROUTINE ======================== github.com/Triztian/pyramid.scanReader in /Users/Tristian/go/src/github.com/Triztian/pyramid/pyramid.go
      30ms      1.02s (flat, cum) 91.89% of Total
         .          .     78:	lsc := bufio.NewScanner(r)
         .          .     79:	lineBuf := bytes.NewBuffer([]byte{})
         .          .     80:	for lsc.Scan() {
         .          .     81:		var lv []int
         .          .     82:
         .       20ms     83:		lineBuf.Write(lsc.Bytes())
         .          .     84:		nsc := bufio.NewScanner(lineBuf)
         .          .     85:		nsc.Split(bufio.ScanWords)
         .      330ms     86:		for nsc.Scan() {
      30ms      250ms     87:			n, err := strconv.ParseInt(string(nsc.Bytes()), 10, 32)
         .          .     88:			if err != nil {
         .          .     89:				return nil, err
         .          .     90:			}
         .      400ms     91:			lv = append(lv, int(n))
         .          .     92:		}
         .          .     93:		lineBuf.Reset()
         .          .     94:
         .       20ms     95:		ns = append(ns, lv)
         .          .     96:	}
         .          .     97:
         .          .     98:	return ns, nil
         .          .     99:}
         .          .    100:
```

## `parse`

```
 λ pyramid ./pyramid.test -test.bench=Parse -test.cpuprofile=cpu.out
goos: darwin
goarch: amd64
pkg: github.com/Triztian/pyramid
BenchmarkParse-4   	     300	   5803607 ns/op
PASS
 λ pyramid go tool pprof -list='parse' cpu.out
Total: 2.39s
ROUTINE ======================== github.com/Triztian/pyramid.parse in /Users/Tristian/go/src/github.com/Triztian/pyramid/pyramid.go
      30ms      1.87s (flat, cum) 78.24% of Total
         .          .     55:
         .          .     56:func parse(s string) ([][]int, error) {
         .          .     57:	ns := [][]int{}
         .          .     58:	var lv int
         .          .     59:	for _, line := range strings.Split(s, "\n") {
         .       20ms     60:		ns = append(ns, []int{})
         .       10ms     61:		line = strings.Trim(line, " ")
         .      870ms     62:		for _, item := range strings.Split(line, " ") {
         .       90ms     63:			n, err := strconv.ParseInt(item, 10, 64)
      10ms       10ms     64:			if err != nil {
         .          .     65:				return nil, err
         .          .     66:			}
      20ms      870ms     67:			ns[lv] = append(ns[lv], int(n))
         .          .     68:		}
         .          .     69:		lv++
         .          .     70:	}
         .          .     71:
         .          .     72:	return ns, nil
 λ pyramid
 ```
