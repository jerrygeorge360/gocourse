package intermediate

import "fmt"


func main(){
	sequence := adder()
	fmt.Print(sequence())

}

func adder() func() int{
	i := 0
	fmt.Println("previous value of i",i)
	return func() int {
		i++
		fmt.Println("added 1 to 1")
		return i
	}

}