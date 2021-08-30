package main

import "strconv"

type calculator struct {
	number1 int
	number2 int
}

func (calculator calculator) sum() string {
	var result = calculator.number1 + calculator.number2
	var op = strconv.Itoa(calculator.number1) + " + " + strconv.Itoa(calculator.number2) + " = " + strconv.Itoa(result)
	return op
}

func (calculator calculator) subtract() string {
	var result = calculator.number1 - calculator.number2
	var op = strconv.Itoa(calculator.number1) + " - " + strconv.Itoa(calculator.number2) + " = " + strconv.Itoa(result)
	return op
}

func (calculator calculator) multiply() string {
	var result = calculator.number1 * calculator.number2
	var op = strconv.Itoa(calculator.number1) + " x " + strconv.Itoa(calculator.number2) + " = " + strconv.Itoa(result)
	return op
}

func (calculator calculator) division() string {
	var result = calculator.number1 / calculator.number2
	var op = strconv.Itoa(calculator.number1) + " / " + strconv.Itoa(calculator.number2) + " = " + strconv.Itoa(result)
	return op
}
