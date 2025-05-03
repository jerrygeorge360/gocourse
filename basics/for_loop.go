package basics

import "fmt"

func main() {

	// for i:=1; i<= 5;i++{
	// 	fmt.Println(i)
	// }

	// numbers := [] int{1,2,3,4,5,6}
	// for index,value:= range numbers{
	// 	fmt.Printf("index: %d,Value: %d\n",index,value)
	// }

	// for i := 1; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd number is ", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	for i := range 10 {
		fmt.Println(10 - i)
	}

}
