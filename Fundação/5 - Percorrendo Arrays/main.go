package main

// percorrendo arrays

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool = true
	c int
	d string
	e float64
)

func main() {
	var meuArray [5]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	meuArray[3] = 40
	meuArray[4] = 50

	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}

	fmt.Println(meuArray[len(meuArray)-1])
}
