package main

import "fmt"

func sum(a, b float64) float64 {
	return a + b
}

func difference(a, b float64) float64 {
	return a - b
}

func main() {

	var num1, num2 float64
	fmt.Print("Введите два числа с плавающей запятой через пробел: ")
	fmt.Scan(&num1, &num2)

	resultSum := sum(num1, num2)
	resultDiff := difference(num1, num2)

	fmt.Printf("Сумма: %.2f + %.2f = %.2f\n", num1, num2, resultSum)
	fmt.Printf("Разность: %.2f - %.2f = %.2f\n", num1, num2, resultDiff)
}
