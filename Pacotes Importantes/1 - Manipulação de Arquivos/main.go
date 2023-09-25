package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Hello, World!")
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	f.Close()

	// Leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	/* Duvida: Se eu tenho 100mb de memória, como ler um arquivo de 1gb?
	   Resposta: Lendo pedaços por pedaços no arquivo, le um pedaço, faz alguma coisa, depois le outro e assim por diante
	*/

	// Leitura de pouco em pouco abrindo o arquivo
	// Utilizar pacote BufIO do Go para trabalhar com buffer e o arquivo ser lido pouco a pouco
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		// Esta convertendo para string o buffer que esta lendo e juntando em um slice o conteudo que esta pegando
		fmt.Println(string(buffer[:n]))
	}

	arquivo2.Close()

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
