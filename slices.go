package main

func main() {
	var numbers[]int
	var numbers1 = []int{1,2,3}
	numbers2 := []int{1}
	slice := make([]int,5)


	fmt.Println("Care to know the capacity", cap(slice))

	dummy := &slice

	fmt.Println(*dummy[2:5]) // can you guess what's contained in the array now?
}
