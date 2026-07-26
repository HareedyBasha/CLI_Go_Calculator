package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func formatString(numbers *[]float64, operators *[]string) (calculationString string) {
	for i, operator := range *operators {
		calculationString += strconv.FormatFloat((*numbers)[i], 'f', -1, 64) + " " + operator + " "
	}
	calculationString += strconv.FormatFloat((*numbers)[len(*numbers)-1], 'f', -1, 64)
	return
}

func calculate(input string) (calculation string) {
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

func getBetweenParentheses(numbers *[]float64, operators *[]string) (equationString string, firstIndex int, secondIndex int, found bool, err string) {
	foundOpeningParens := false
	foundClosingParens := false

	var openingParensIndex int
	var closingParensIndex int

	for i, operator := range *operators {
		if operator == "(" {
			if foundOpeningParens {
				err = "Syntax Error"
				return
			}
			foundOpeningParens = true
			openingParensIndex = i
		}
		if operator == ")" {
			if foundOpeningParens {
				foundClosingParens = true
				closingParensIndex = i
				break
			} else {
				err = "Syntax Error"
				return
			}
		}
	}

	if foundOpeningParens && foundClosingParens {
		if openingParensIndex == closingParensIndex+1 {
			err = "Syntax Error"
			return
		}

		found = true
		numbersBetweenParens := (*numbers)[openingParensIndex:closingParensIndex]
		operatorsBetweenParens := (*operators)[openingParensIndex+1 : closingParensIndex]

		equationString = formatString(&numbersBetweenParens, &operatorsBetweenParens)
	} else if foundOpeningParens {
		err = "Syntax Error"
	}

	return
}

func applyPrecedence(precendents map[string]func(num1, num2 float64) float64, nums *[]float64, operators *[]string) {

	foundPrecendent := true

	for foundPrecendent {
		foundPrecendent = false
		for i, operator := range *operators {
			var newNum float64
			if fn, exists := precendents[operator]; exists {
				newNum = fn((*nums)[i], (*nums)[i+1])

				// Gets a slice of num up to but not including the first number of the two numbers of the operator
				// Appends to that slice the result of appending to newNum a slice of num that starts after the second number of the two numbers of the operator till the end
				// So nums will have the two numbers of the operator replaced by the result of their calculation using their operator
				*nums = append((*nums)[:i], append([]float64{newNum}, (*nums)[i+2:]...)...)

				// Makes operators the result of appending all of the operators after operator[i] to all of the operators before operator[i]
				// so operators will have operator[i] removed
				*operators = append((*operators)[:i], (*operators)[i+1:]...)

				foundPrecendent = true
				break
			}

		}
	}

}

func advancedCalculate(input string) (formattedString string, calculationErr string, result float64) {
	var numbers []float64
	var operators []string

	// a map to check whether an operator is valid or not in O(1) time, the value type is struct because an empty struct takes up 0 bytes
	validOperators := map[string]struct{}{"*": {}, "/": {}, "+": {}, "-": {}}

	multipOperators := map[string]func(num1, num2 float64) float64{"*": func(num1, num2 float64) float64 { return num1 * num2 },
		"/": func(num1, num2 float64) float64 { return num1 / num2 }}

	additiveOperators := map[string]func(num1, num2 float64) float64{"+": func(num1, num2 float64) float64 { return num1 + num2 },
		"-": func(num1, num2 float64) float64 { return num1 - num2 }}

	invalidRegex := regexp.MustCompile(`[^0-9\+\-\/\*\(\)\s]`)
	regex := regexp.MustCompile(`\d+|[\*\/\+\-\)\(]`)

	if invalidRegex.MatchString(input) {
		calculationErr = "Syntax Error"
		return
	}

	tokens := regex.FindAllString(input, -1)

	if len(tokens)%2 == 0 {
		calculationErr = "Syntax Error"
		return
	}

	for i, token := range tokens {
		cleanToken := strings.TrimSpace(token)
		switch i % 2 {
		case 0:
			num, err := strconv.ParseFloat(cleanToken, 64)
			if err != nil {
				calculationErr = "Syntax Error"
				return
			}
			numbers = append(numbers, num)
		case 1:
			if _, exists := validOperators[cleanToken]; exists {
				operators = append(operators, cleanToken)
			} else {
				calculationErr = "Syntax Error"
				return
			}
		}
	}

	originalNumbers := make([]float64, len(numbers))
	originalOperators := make([]string, len(operators))

	copy(originalNumbers, numbers)
	copy(originalOperators, operators)

	// foundParens := true
	// var equationString string
	// var firstIndex, secondIndex int

	// for foundParens {
	// 	equationString, firstIndex, secondIndex, foundParens, calculationErr = getBetweenParentheses(&numbers, &operators)

	// 	if calculationErr != "" {
	// 		return
	// 	}

	// 	if foundParens {
	// 		_, _, calculationBetweenParensResult := advancedCalculate(equationString)
	// 		// Gets a slice of num up to but not including the first number of the equation numbers of the parentheses equation
	// 		// Appends to that slice the result of appending to the calculation result of the calculation result a slice of num that
	// 		// starts after the last number of the numbers of the parentheses equation till the end
	// 		// So nums will have the numbers of the parentheses equation replaced by the result of their calculation
	// 		numbers = append(numbers[:firstIndex], append([]float64{calculationBetweenParensResult}, numbers[secondIndex:]...)...)

	// 		// Makes operators the result of appending all of the operators after operator[secondIndex] to all of the operators before operator[firstIndex]
	// 		// so operators will have the parentheses and the operators within it removed
	// 		operators = append(operators[:firstIndex], operators[secondIndex+1:]...)
	// 	}
	// }

	applyPrecedence(multipOperators, &numbers, &operators)
	applyPrecedence(additiveOperators, &numbers, &operators)

	result = numbers[0]

	if math.IsInf(result, 0) || math.IsNaN(result) {
		calculationErr = "Math Error"
		return
	}

	calculationString := formatString(&originalNumbers, &originalOperators)

	calculationString += " = " + strconv.FormatFloat(result, 'f', -1, 64)

	timeOfCalculation := time.Now().Format("03:04:05 PM") + ": "
	formattedString = timeOfCalculation + calculationString

	return

}

func calculationLoop(history *[]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a simple mathematical in the following style:- \n {num} {operator (*, /, +, -)} {num}, space separated \n type 'menu' to return to the menu")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		if strings.TrimSpace(strings.ToLower(input)) == "menu" {
			return
		}
		calculationResult := calculate(input)

		if calculationResult != "" {
			*history = append(*history, calculationResult)
		}
	}
}

func advancedCalculationLoop(history *[]string) (quit bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a mathematical equation using only the *,/,+, and - operators\ntype 'menu' to return to the menu or type 'exit' to quit the program")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		if strings.TrimSpace(strings.ToLower(input)) == "menu" {
			return
		}

		if strings.TrimSpace(strings.ToLower(input)) == "exit" {
			return true
		}

		formattedString, calculationErr, calculationResult := advancedCalculate(input)

		if calculationErr == "" {
			*history = append(*history, formattedString)
			fmt.Println(calculationResult)

		} else {
			fmt.Println(calculationErr)
		}
	}

}

func EasyCalculator() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input a simple mathematical in the following style:- \n {num} {operator (*, /, +, -)} {num}, space separated \n type 'exit' to exit the program")

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		if strings.TrimSpace(strings.ToLower(input)) == "exit" {
			return
		}

		calculate(input)
	}
}
func IntermediateCalculator() {
	reader := bufio.NewReader(os.Stdin)
	var history []string

	fmt.Println("Welcome to the \033[1mGo CLI Calculator\033[0m, please select an option from below by typing in the corresponding number or name")
	for {
		fmt.Println("(1) \033[1mCalculate\033[0m\n(2) \033[1mShow History\033[0m\n(3) \033[1mClear History\033[0m\033[0m\n(4) \033[1mExit\033[0m")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		cleanInput := strings.TrimSpace(strings.ToLower(input))

		switch cleanInput {
		case "1", "calculate":
			calculationLoop(&history)
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

func AdvancedCalculator() {
	reader := bufio.NewReader(os.Stdin)
	var history []string

	fmt.Println("Welcome to the \033[1mGo CLI Calculator\033[0m, please select an option from below by typing in the corresponding number or name")
	for {
		fmt.Println("(1) \033[1mCalculate\033[0m\n(2) \033[1mShow History\033[0m\n(3) \033[1mClear History\033[0m\033[0m\n(4) \033[1mExit\033[0m")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		cleanInput := strings.TrimSpace(strings.ToLower(input))

		switch cleanInput {
		case "1", "calculate":
			if advancedCalculationLoop(&history) {
				return
			}
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
	AdvancedCalculator()
}
