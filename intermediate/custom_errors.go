package intermediate

import "fmt"


func main(){
	err := doSomething()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Operation completed successfully")
}

type customError struct{
	code int
	message string
}
// Error returns the error message implementing Error() method of error interface

func(e *customError) Error() string{
	return fmt.Sprintf("Error %d:%s",e.code,e.message)
}

// Function that returns a custom error
func doSomething()error{
	return &customError{
		code: 500,
		message:"Something went wrong",
	}
}