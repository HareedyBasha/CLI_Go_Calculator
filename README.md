# Calculator Project

## 📋 Overview
This project is a **command-line calculator** **Go programming**. It progresses through three levels—**easy to intermediate**—to build familiarity with Go syntax, fundamental programming concepts, and data structures. Each level introduces new features and Go concepts, ensuring a structured learning path for beginners.

## 🎯 Objective
Develop a command-line calculator that performs mathematical operations, helping trainees master:
- Go syntax (variables, functions, control structures).
- Data structures (slices, structs, maps).
- String parsing and algorithmic thinking.

---

## 🛠️ Project Levels and Requirements

### Level 1: Easy - Basic Calculator
**Goal**: Familiarize trainees with Go syntax, basic input/output, and simple functions.

**Requirements**:
- Support basic arithmetic operations: **addition (+)**, **subtraction (-)**, **multiplication (*)**, and **division (/)**.
- Accept user input for two numbers (`float64`) and an operator via the command line.
- Display the result of the operation.
- Handle **division by zero** with an error message.
- Provide an option to exit the program.

**Features**:
1. Prompt users to enter two numbers and an operator (+, -, *, /).
2. Perform the selected operation and show the result.
3. Display errors for invalid operators or division by zero.
4. Allow users to exit the program.

**Go Concepts Covered**:
- Variables and data types (`float64`, `string`).
- Standard input/output (e.g., `fmt.Scanln` or `bufio.Reader`).
- Functions to modularize operations (e.g., `add`, `subtract`, `multiply`, `divide`).
- Conditionals (`if-else` or `switch`) for operator selection.
- Basic error handling with return values or `panic` (optional).

---

### Level 2: Easy-Intermediate - Enhanced Calculator with History
**Goal**: Introduce slices, structs, and menu-driven programming.

**Requirements**:
- Extend the basic calculator to store a **history of calculations**.
- Allow users to view all calculations performed in the session.
- Use a **slice** to store calculation records (e.g., "2 + 3 = 5").
- Provide a menu to select: perform a calculation, view history, clear history, or exit.
- Handle invalid menu inputs gracefully.

**Features**:
1. Store each calculation (numbers, operator, result) in a slice.
2. Display a menu with options: **(1) Calculate**, **(2) View History**, **(3) Clear History**, **(4) Exit**.
3. Show history with formatted output (e.g., index, expression, result).
4. Allow users to clear the history.
5. Handle invalid menu inputs with clear error messages.

**Go Concepts Covered**:
- Slices (`[]string` or custom struct for history).
- Structs to represent calculations (e.g., `struct { num1, num2 float64; op string; result float64 }`).
- Loops (`for`, `range`) to iterate over history.
- String formatting with `fmt.Sprintf` for history display.
- Menu-driven program flow with `switch`.

---

### Level 3: Intermediate - Advanced Calculator with Expressions
**Goal**: Introduce maps, string parsing, and complex logic for evaluating mathematical expressions.

**Requirements**:
- Allow users to input a full mathematical expression (e.g., "2 + 3 * 4").
- Support **operator precedence** (e.g., * and / before + and -).
- Use a **map** to store operator functions or precedence levels.
- Handle invalid expressions with clear error messages.
- Maintain the history feature from Level 2.

**Features**:
1. Parse a string input (e.g., "2 + 3 * 4") into tokens (numbers and operators).
2. Evaluate expressions respecting operator precedence (e.g., compute 3 * 4 first, then add 2).
3. Store the expression and result in history.
4. Display errors for invalid expressions (e.g., "2 + + 3" or "5 / 0").
5. Allow users to re-run a previous expression from history.

**Go Concepts Covered**:
- Maps (e.g., `map[string]int` for operator precedence or `map[string]func(float64, float64) float64` for operations).
- String parsing (using `strings.Split` or custom tokenization).
- Stack-based expression evaluation (using slices as stacks for numbers and operators).
- Error handling with custom error types.
- Pointers (optional, for advanced history manipulation).

---

## ℹ️ Additional Notes
- **Code Organization**: Encourage trainees to use separate files for different components (e.g., `operations.go` for calculation logic, `history.go` for history management).
- **Testing**: Introduce unit testing with Go’s `testing` package to verify operations and expression evaluation.
- **Extensions** (for advanced learners):
  - Support parentheses in expressions (requires a more complex parser).
  - Implement persistent history using file I/O (`os` or `io` packages).
  - Add support for additional operations (e.g., modulus, exponentiation).
- **Resources**:
  - [Official Go Tour](https://tour.golang.org) for learning syntax.
  - [Go by Example](https://gobyexample.com) for practical code snippets.
  - Practice problems on [Exercism](https://exercism.org) or [LeetCode](https://leetcode.com) to reinforce concepts.