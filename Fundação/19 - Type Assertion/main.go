package main

import "fmt"

// Type assertion e uma forma de converter um tipo de variavel para outro tipo de variavel
func main() {
	var minhaVar interface{} = "Marcelo Martins"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res e %v e o valor de ok e %v\n", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 e %v\n", res2)
}
