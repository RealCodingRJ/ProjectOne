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

type Option struct {
	message string
}

func print_message(message Option) {
	fmt.Println(message.message)
}

func get_i() {

	var num int

	n := Number{number: num}

	fmt.Scanln(&n.number)

	get_mod_number(n.number)
}

func main() {

	num_message := Option{
		message: "Enter a Number: ",
	}

	o_message := Option{
		message: "[WELCOME TO EVEN Or Odd Application]",
	}

	print_message(o_message)

	print_message(num_message)
	go get_i()
	get_i()
}
