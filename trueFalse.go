package main

import (
	"fmt"
	"strconv"
)

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessge := ""
	secretValue := 88

	//Insert your code from here
	if value == secretValue {
		resultMessge = "Well Done! Your guess is correct"
	}

	if value < secretValue {
		resultMessge = "Too low, try again next time!"
	}

	if value > secretValue {
		resultMessge = "Too high, try again next time!"
	}

	fmt.Println(resultMessge)

	//do not remove this line
	return resultMessge
}

func main() {

	// Taking input as String from user
	var input string
	var guess int

	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&input)

	// Convert string to integer. Checking integer during input will only check first entered character, the rest of the string will result in command not found: error.
	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter integer value only.")
	} else {
		compare(guess)
	}

}
