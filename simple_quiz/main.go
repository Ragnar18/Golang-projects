package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")
	var name string

	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")

	} else {
		fmt.Println("You can't play!")
		return
	}

	score := 0
	num_questions := 2

	fmt.Printf("What GPU is better, the RTX 3080 or RTX 3090? ")
	var answer int
	var answer2 int
	fmt.Scan(&answer)

	fmt.Println(answer)

	if answer == 3090 {
		fmt.Println("Correct!")
		score += 1

	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	fmt.Scan(&answer2)

	if answer2 == 12 {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("Your score:%v out of %v questions", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Println("You scored %v%%. ", percent)

}
