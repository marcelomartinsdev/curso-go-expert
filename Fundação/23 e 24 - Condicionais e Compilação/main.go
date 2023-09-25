package main

func main() {
	a := 1
	b := 2
	c := 3

	// Nao tem else if no Go
	if a > b {
		println(a)
	} else {
		println(b)
	}

	if a > b && c > a {
		println("a > b && c > a")
	}

	if a > b || c > a {
		println("a > b && c > a")
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}
}

// 24 - Para compilar um código em Go, utilizar comando go build, caso não tenha um módulo criado pode-se usar go build main.go, irá gerar o binário main.exe
