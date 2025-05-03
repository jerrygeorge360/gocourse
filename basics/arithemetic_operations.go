package basics

import "fmt"

func main() {
	var a, b int = 10, 30
	var result int
	result = a + b
	fmt.Println(result)

	const p float64 = 22/7 // evaluates from right to left hence prints an integer
	fmt.Println(p)


}
