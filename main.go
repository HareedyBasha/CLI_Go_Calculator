package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EasyCalculator() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input a simple mathematical in the following style:- \n {num} {operator (*, /, +, -)} {num}, space separated \n type 'exit' to exit the program")

loop:
	for true {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		if strings.TrimSpace(strings.ToLower(input)) == "exit" {
			return
		}

		var num1, num2 *float64
		var operator *string

		tokens := strings.Fields(input)

		for i, token := range tokens {
			cleanToken := strings.TrimSpace(token)
			switch i {
			case 0:
				num1Val, err := strconv.ParseFloat(cleanToken, 64)
				if err != nil {
					fmt.Printf("Error at the first number: %v\n", err)
					continue loop
				}
				num1 = &num1Val
			case 1:
				operator = &cleanToken
			case 2:
				num2Val, err := strconv.ParseFloat(cleanToken, 64)
				if err != nil {
					fmt.Printf("Error at the second number number: %v\n", err)
					continue loop
				}
				num2 = &num2Val
			default:
				fmt.Println("Error: you haven't followed the formula of (num operator num) and may have written more that allowed, please try again")
				continue loop
			}
		}

		if num1 == nil {
			fmt.Println("Error: The first number was not provided")
			continue
		} else if operator == nil {
			fmt.Println("Error: A mathematical operator was not provided was not provided")
			continue
		} else if num2 == nil {
			fmt.Println("Error: The second number was not provided")
			continue
		}

		switch *operator {
		case "*":
			fmt.Println(*num1 * *num2)
		case "/":
			if *num2 == 0 {
				fmt.Println("Math Error: Cannot divide by zero")
			} else {
				fmt.Println(*num1 / *num2)
			}
		case "+":
			fmt.Println(*num1 + *num2)
		case "-":
			fmt.Println(*num1 - *num2)
		default:
			fmt.Println("Error: The provided string was not a valid mathematical operator, please use one of the following (*, /, +, -)")
		}

	}
}

func main() {
	EasyCalculator()
}
