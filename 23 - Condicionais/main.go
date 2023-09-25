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
