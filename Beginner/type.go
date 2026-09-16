package main

import (
	"fmt"
)

var (
	a bool       = true
	b string     = "Hello"
	c int        = 42
	d int8       = 127
	e int16      = 32767
	f int32      = 2147483647
	g int64      = 9223372036854775807
	h uint       = 42
	i uint8      = 255
	j uint16     = 65535
	k uint32     = 4294967295
	l uint64     = 18446744073709551615
	m uintptr    = 0xFFFFFFFFFFFFFFFF
	n float32    = 3.14
	o float64    = 3.141592653589793
	p complex64  = complex(1, 2)
	q complex128 = complex(1, 2)
	r rune       = 'a'
	s byte       = 255
	t int
)

func main() {

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of c: %T\n", c)
	fmt.Printf("Type of d: %T\n", d)
	fmt.Printf("Type of e: %T\n", e)
	fmt.Printf("Type of f: %T\n", f)
	fmt.Printf("Type of g: %T\n", g)
	fmt.Printf("Type of h: %T\n", h)
	fmt.Printf("Type of i: %T\n", i)
	fmt.Printf("Type of j: %T\n", j)
	fmt.Printf("Type of k: %T\n", k)
	fmt.Printf("Type of l: %T\n", l)
	fmt.Printf("Type of m: %T\n", m)
	fmt.Printf("Type of n: %T\n", n)
	fmt.Printf("Type of o: %T\n", o)
	fmt.Printf("Type of p: %T\n", p)
	fmt.Printf("Type of q: %T\n", q)
	fmt.Printf("Type of r: %T\n", r)
	fmt.Printf("Type of s: %T\n", s)
	fmt.Printf("Type of t: %T\n", t)
	fmt.Printf("default value of t: %v", t)
}
