package main

import (
	"fmt"
)

func main() {
	fmt.Println("Дебильный калькулятор")
	fmt.Println("Какое действие Вы хотите выполнить?, (+, -, /, * )")

	var a float64
	var b float64
	var action string

	fmt.Scan(&action)

	fmt.Println("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Println("Введите второе число: ")
	fmt.Scan(&b)

	switch {
	case action == "+":
		fmt.Println("Сумма чисел " + fmt.Sprint(a+b))
	case action == "-":
		fmt.Println("Разность чисел " + fmt.Sprint(a-b))

	case action == "/":
		fmt.Println("Деление чисел " + fmt.Sprint(a/b))

	case action == "*":
		fmt.Println("Произведение чисел " + fmt.Sprint(a*b))

	}
	fmt.Scan(&a)
}
