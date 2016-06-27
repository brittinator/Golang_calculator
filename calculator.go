package main

import (
	"fmt"
	"strconv"
	// "math"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello there! So you want to use this calculator?")
	operator()
}

func operator() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What operation would you like to perform?")
	operation, _ := reader.ReadString('\n')
	op := strings.TrimSpace(operation)
	if op == "+" || op == "add" {
		fmt.Println(op)
		nums()
	} else if op == "-" || op == "sub" {
		first, second := num()
	} else if op == "*" || op == "mul" {
	} else if op == "/" || op == "div" {
	} else {
		fmt.Println("I'm not sure what you mean. Exiting")
	}
}

func nums() (int, int) {
	// Asks for 1st and 2nd number, trims whitespace then converts to a number
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your first number: ")
	firStr, _ := reader.ReadString('\n')
	// firInt := strings.TrimSpace(firStr)
	firInt, _ := strconv.Atoi(firInt)
	fmt.Println("Please enter your second number: ")
	sec, _ := reader.ReadString('\n')
	return firInt, secInt
}
