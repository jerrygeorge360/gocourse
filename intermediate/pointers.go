package intermediate

import "fmt"

type CustomInt int

func (n *CustomInt) multiplyBySeven() CustomInt {
	return *n * 7;
}

func main(){
	var ptr *CustomInt
	var a CustomInt = 10
	ptr = &a

	fmt.Println(a)
	fmt.Println(ptr)

	
	fmt.Println(ptr.multiplyBySeven())

}
