package math

import "fmt"

func Fibonacci(n int) {
	FmtFibonacci(n, "%d\n")
}

func FmtFibonacci(n int, fmtString string) {
	if n == 0 {
		return
	}

	var currentNumber, prevNumber int = 1, 1

	fmt.Printf(fmtString, currentNumber)

	if n > 1 {
		fmt.Printf(fmtString, currentNumber)
	}

	for i := 3; i <= n; i++ {
		var temp int = currentNumber
		currentNumber += prevNumber

		fmt.Printf(fmtString, currentNumber)

		prevNumber = temp
	}

}
