package main

import "fmt"

func get_mod_number(a int) {

	if a%2 == 0 {
		fmt.Println("Even")
	}

	if a%2 != 0 {

		fmt.Println("Odd")
	}

}

type Number struct {
	number int
}

func get_i() {

	var num int

	n := Number{number: num}

	fmt.Scanln(&n.number)

	get_mod_number(n.number)
}

func main() {

	fmt.Println("Enter Num: ")
	go get_i()
	get_i()
}
