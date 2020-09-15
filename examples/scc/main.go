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

	n, m := getInt(), getInt()

	g := acl.SccGraph(n)
	for i := 0; i < m; i++ {
		u, v := getInt(), getInt()
		g.AddEdge(u, v)
	}

	scc := g.Scc()

	out(len(scc))
	for _, e := range scc {
		fmt.Print(len(e))
		for _, x := range e {
			fmt.Print(" ", x)
		}
		out()
	}
}
