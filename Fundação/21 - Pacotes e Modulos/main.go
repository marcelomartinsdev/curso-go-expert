package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/marcelomartinsdev/curso-go-expert/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v\n", s)
	fmt.Println(uuid.New())
}
