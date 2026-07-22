package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func EasyCalculator() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input a simple mathematical in the following style:- \n {num} {operator (*, /, +, -)} {num}, space separated \n type 'exit' to exit the program")

	for true {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		if strings.TrimSpace(strings.ToLower(input)) == "exit" {
			return
		}

		Calculate(input)
	}
}

func Calculate(input string) (calculation string) {
	var num1, num2 *float64
	var operator *string
	var result float64

	tokens := strings.Fields(input)

	for i, token := range tokens {
		cleanToken := strings.TrimSpace(token)
		switch i {
		case 0:
			num1Val, err := strconv.ParseFloat(cleanToken, 64)
			if err != nil {
				fmt.Printf("Error at the first number: %v\n", err)
				return
			}
			num1 = &num1Val
		case 1:
			operator = &cleanToken
		case 2:
			num2Val, err := strconv.ParseFloat(cleanToken, 64)
			if err != nil {
				fmt.Printf("Error at the second number number: %v\n", err)
				return
			}
			num2 = &num2Val
		default:
			fmt.Println("Error: you haven't followed the formula of (num operator num) and may have written more that allowed, please try again")
			return
		}
	}

	if num1 == nil {
		fmt.Println("Error: The first number was not provided")
		return
	} else if operator == nil {
		fmt.Println("Error: A mathematical operator was not provided was not provided")
		return
	} else if num2 == nil {
		fmt.Println("Error: The second number was not provided")
		return
	}

	switch *operator {
	case "*":
		result = *num1 * *num2
	case "/":
		if *num2 == 0 {
			fmt.Println("Math Error: Cannot divide by zero")
			return
		} else {

			result = *num1 / *num2
		}
	case "+":
		result = *num1 + *num2
	case "-":
		result = *num1 - *num2
	default:
		fmt.Println("Error: The provided string was not a valid mathematical operator, please use one of the following (*, /, +, -)")
		return
	}

	fmt.Println(result)

	operation := strconv.FormatFloat(*num1, 'f', -1, 64) + " " + *operator + " " + strconv.FormatFloat(*num2, 'f', -1, 64) + " = " + strconv.FormatFloat(result, 'f', -1, 64)
	timeOfCalculation := time.Now().Format("03:04:05 PM") + ": "
	calculation = timeOfCalculation + operation
	return

}

func CalculationLoop(history *[]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a simple mathematical in the following style:- \n {num} {operator (*, /, +, -)} {num}, space separated \n type 'menu' to return to the menu")
	for true {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		if strings.TrimSpace(strings.ToLower(input)) == "menu" {
			return
		}
		calculationResult := Calculate(input)

		if calculationResult != "" {
			*history = append(*history, calculationResult)
		}
	}
}

func IntermediateCalculator() {
	reader := bufio.NewReader(os.Stdin)
	var history []string

	fmt.Println("Welcome to the \033[1mGo CLI Calculator\033[0m, please select an option from below by typing in the corresponding number or name")
	for true {
		fmt.Println("(1) \033[1mCalculate\033[0m\n(2) \033[1mShow History\033[0m\n(3) \033[1mClear History\033[0m\033[0m\n(4) \033[1mExit\033[0m")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		cleanInput := strings.TrimSpace(strings.ToLower(input))

		switch cleanInput {
		case "1", "calculate":
			CalculationLoop(&history)
		case "2", "show history":
			if len(history) == 0 {
				fmt.Print("\n\033[1mHistory is empty\033[0m\n\n")
			} else {
				fmt.Println("\n\033[1m===== History =====\033[0m")
				for _, entry := range history {
					fmt.Println(entry)
				}
				fmt.Print("\033[1m===== History =====\n\n\033[0m")
			}
		case "3", "clear history":
			history = history[:0] // history[:0] means that the new slice will read from the beginning up to but not including index 0
			//, and since there is nothing between index 0 and the beginning
			//, the result is an empty slice with a length of zer0
			fmt.Print("\n\033[1mHistory cleared\033[0m\n\n")
		case "4", "exit":
			return
		default:
			fmt.Println("Error: Please input a correct option")
			continue
		}

	}

}

func main() {
	IntermediateCalculator()
}
