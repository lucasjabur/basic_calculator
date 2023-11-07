package main

import (
	"fmt"
)

func main() {

	var fst_number, snd_number float64
	var operator string

	fmt.Printf("Enter with the first value: ")
	fmt.Scanln(&fst_number)

	fmt.Printf("Enter with the operator (+ - * /): ")
	fmt.Scanln(&operator)

	fmt.Printf("Enter with the second value: ")
	fmt.Scanln(&snd_number)

	switch operator {

	case "+":
		fmt.Printf("%.3f %s %.3f = %.3f\n", fst_number, operator, snd_number, fst_number+snd_number)

	case "-":
		fmt.Printf("%.3f %s %.3f = %.3f\n", fst_number, operator, snd_number, fst_number-snd_number)

	case "*":
		fmt.Printf("%.3f %s %.3f = %.3f\n", fst_number, operator, snd_number, fst_number*snd_number)

	case "/":

		if snd_number == 0.0 {
			fmt.Println("Not possible to divide by zero!")
		} else {
			fmt.Printf("%.3f %s %.3f = %.3f\n", fst_number, operator, snd_number, fst_number/snd_number)
		}

	default:
		fmt.Println("Invalid operator!")

	}

}
