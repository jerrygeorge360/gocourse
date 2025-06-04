package main

import (
	"fmt"
	"strconv"
)


func main(){

	numStr := "12345"
	num,err := strconv.Atoi(numStr)
	if err != nil{
		fmt.Println("Error parsing the value: ",err)
	
	}
	fmt.Println("Parsed Intergr:",num)
	fmt.Println("Parsed INterger",num+1)

	numistr,err := strconv.ParseInt(numStr,10,32)
	if err != nil{
		fmt.Println("Error parsing the value: ",err)
	
	}
	fmt.Println("Parsed Integer",numistr)


	binaryStr := "10111"
	decimal,err := strconv.ParseInt(binaryStr,2,64)
	if err != nil{
		fmt.Println("Error parsing the value: ",err)
	
	}
	fmt.Println("Parsed binary to decimal",decimal)
}

// TODO : study time in go