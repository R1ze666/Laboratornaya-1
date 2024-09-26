package main

import "fmt"

func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {

	var num1, num2, num3 float64

	fmt.Print("Введите три числа (целые или с плавающей запятой) через пробел: ")
	fmt.Scan(&num1, &num2, &num3)

	result := average(num1, num2, num3)

	fmt.Printf("Среднее арифметическое чисел %.2f, %.2f, %.2f = %.2f\n", num1, num2, num3, result)
}
