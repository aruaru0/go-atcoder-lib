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

	d := acl.NewDsu(n)
	for i := 0; i < q; i++ {
		t, u, v := getInt(), getInt(), getInt()
		if t == 0 {
			d.Merge(u, v)
		} else {
			if d.Same(u, v) {
				out("1")
			} else {
				out(0)
			}
		}
	}
}
