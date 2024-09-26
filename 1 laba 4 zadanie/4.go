package main

import "fmt"

func main() {

	var num1, num2 int
	fmt.Print("Введите два числа через пробел: ")
	fmt.Scan(&num1, &num2)

	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2

	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
	fmt.Printf("%d - %d = %d\n", num1, num2, difference)
	fmt.Printf("%d * %d = %d\n", num1, num2, product)

	if num2 != 0 {
		quotient := float64(num1) / float64(num2)
		fmt.Printf("%d / %d = %.2f\n", num1, num2, quotient)
	} else {
		fmt.Println("Деление на ноль невозможно!")
	}
}
