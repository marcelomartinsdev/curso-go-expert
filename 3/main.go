package main

// criação de tipos e variáveis com declaração curta

const a = "Hello, World!"

type ID int

var (
	b bool = true
	c int
	d string
	e float64
	f complex128
	g rune
	h byte
	i ID = 1
)

func main() {
	a := "X"
	println(a, b, c, d, e, f, g, h, i)
}
