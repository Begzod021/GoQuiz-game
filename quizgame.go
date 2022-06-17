package main
import (
	"fmt"
)


func main() {
	fmt.Println("Welcome to my quiz game!")
	
	fmt.Printf("Enter your name: ")	
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, Welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age int
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("You can play!")
	} else {
		fmt.Println("You can't play ")
		return
	}

	fmt.Println("Continue")

	score := 0
	num_question := 2


	fmt.Printf("What is better, RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer + " " + answer2 == "RTX 3090" || answer + " " + answer2 == "rtx 3090"{
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Incorrect!")
	}


	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++

	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v.\n", score, num_question)
	percent := (float64(score) / float64(num_question)) * 100
	fmt.Printf("You scored: %v%%.", percent)
}