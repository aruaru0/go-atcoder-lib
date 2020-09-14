package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/aruaru0/go-atcoder-lib/acl"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	n, q := getInt(), getInt()

	fw := acl.FenwickTree(n)
	for i := 0; i < n; i++ {
		a := getInt()
		fw.Add(i, a)
	}

	for i := 0; i < q; i++ {
		t := getInt()
		if t == 0 {
			p, x := getInt(), getInt()
			fw.Add(p, x)
		} else {
			l, r := getInt(), getInt()
			out(fw.Sum(l, r))
		}
	}
}
