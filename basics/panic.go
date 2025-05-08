package basics

import "fmt"


func main(){

	// Example of valid input
	// process(10)

	// Invalid input
	process(-3)
}

func process(input int){
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0{
		fmt.Println("Before Panic")
		panic("input must be non-negative number")

	}
	fmt.Println("processing input",input)
}