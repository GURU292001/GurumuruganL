package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 0.0
	b := 0.0
	operator := ""
	fmt.Print("Enter the value a :")
	fmt.Scan(&a)
	fmt.Print("Enter the value b :")
	fmt.Scan(&b)
	fmt.Print("Enter the Operation like '-','/','+','*' : ")
	fmt.Scan(&operator)
	lValue, lerr := Problem1(a, b, operator)
	if lerr != nil {
		fmt.Println(lerr)
	} else {
		fmt.Println("Result :", lValue)
	}

}

func Problem1(a, b float64, operator string) (float64, error) {

	switch operator {
	case "-":
		return a - b, nil
	case "+":
		return a + b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("please provide the valid operator")
	}
}
