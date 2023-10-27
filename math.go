package main
import "fmt"

func main() {
	fmt.Println(SumInt(10, 11))
}

// Sum is a function that returns the sum of two integers
func SumInt(a int, b int) int {
	return a + b
}