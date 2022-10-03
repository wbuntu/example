package main

import (
	"fmt"
)

// Stringy takes a string and returns a possibly different string.
type Stringy func(string) string

// PrintString declares a method on the Stringy function.
// This way we can make our function match the StringPrinter
// interface.
func (s Stringy) PrintString(str string) {
	println(s(str))
}

var a Stringy = func(str string) string {
	return "Got " + str
}

func main() {
	a.PrintString("Test")
	fmt.Printf("func %v %p %p\n", a, a, &a)
	b := a
	b.PrintString("Test")
	fmt.Printf("func %v %p %p\n", b, b, &b)

	d := 28
	fmt.Printf("int %v %p\n", d, &d)
	e := d
	fmt.Printf("int %v %p\n", e, &e)

	f := make(chan byte, 1)
	fmt.Printf("chan %v %p %p\n", f, f, &f)
	g := f
	fmt.Printf("chan %v %p %p\n", g, g, &g)

	h := make([]byte, 1)
	fmt.Printf("slice %v %p %p\n", h, h, &h)
	i := h
	fmt.Printf("slice %v %p %p\n", i, i, &i)

	j := make(map[byte]byte, 1)
	fmt.Printf("map %v %p %p\n", j, j, &j)
	k := j
	fmt.Printf("map %v %p %p\n", k, k, &k)

	var l,m,n,o,p interface{}
	l = b
	fmt.Printf("l:func %v %p %p\n", l, l, &l)
	m = e
	fmt.Printf("m:int %v %p %p\n", m, m, &m)
	n = g
	fmt.Printf("n:chan %v %p %p\n", n, n, &n)
	o = i
	fmt.Printf("o:slice %v %p %p\n", o, o, &o)
	p = j
	fmt.Printf("p:map %v %p %p\n", p, p, &p)
}
