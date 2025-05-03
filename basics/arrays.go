package basics

import "fmt"

func main() {
	var arrayName [10]int
	fruits := [4]string{"apple","banana","peach","mango"}
	fmt.Println(arrayName)
	arrayName[9] = 8
	fmt.Println(arrayName)
	fmt.Println(fruits)
	
	var copied_array *[10]int
	copied_array = &arrayName
	copied_array[0] = 3

	fmt.Println(copied_array)
	fmt.Println(arrayName)

	for i:= 0;i<len(arrayName);i++{
		fmt.Println(arrayName[i])
	}

	for _,v := range arrayName{
		fmt.Println(v)
	}

	matrix := [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	fmt.Println(matrix)

}
