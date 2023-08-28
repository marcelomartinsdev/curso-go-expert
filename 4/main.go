package main

// importação do pacote fmt

import "fmt"

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
	fmt.Printf("O tipo de E é %T\n", i)
}
