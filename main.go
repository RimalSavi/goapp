package main

import (
	"fmt"

	"github.com/RimalSavi/goapp/calculator"
)

func main() {
	sum := calculator.Add(2, 4)

	calc := calculator.Calculator{}

	mult := calc.Multiply(3, 5)

	fmt.Println(sum)
	fmt.Println(mult)
}
