package main

import "fmt"

func main() {
	number := 10

	switch {
	case number > 1:
		fmt.Println("Number is greater than 1")
		fallthrough
	case number < 11:
		fmt.Println("Number < 11")
	}
}
