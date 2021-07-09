package main

import (
	"fmt"
	"github.com/raydeth/gocourse/lesson_02/math"
)

func main() {
	math.Fibonacci(10)
	math.FmtFibonacci(15, "Число: %d\n")

	defer fmt.Println("Finished")
}
