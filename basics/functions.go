package main

import "fmt"


func main(){

fmt.Println(add(4,5))

greet := func(){
	fmt.Println("Hello anonymous function")
}

greet()

}

func add(a,b int) int{

	return a+b
}