package main

func main() {
	a := 12
	b := 23
	c := 22

	if a > b && a > c {
		println("A is the biggest Number!")
	} else if b > a && b > c {
		println("B is the biggest Number!")
	} else {
		println("C is the biggest Number!")
	}
}
