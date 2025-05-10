package intermediate

import (
	"errors"
	"fmt"
)

// creating a function with error functionality
func sqrt(x float64)(float64,error){
	if x<0{
		return 0,errors.New("Math error:square root of negative numbers")
	}
	return 1,nil
}

func process(data []byte) error{
	if len(data) == 0{
		return errors.New("Error: Empty data")
	}
	return nil
}

func main(){
	// result, err := sqrt(-16)

	// // handling the error
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// data := []byte{}
	// process_data := process(data)

	// if process_data != nil{
	// 	fmt.Println("Error",process_data)
	// 	return
	// }
	// fmt.Println("Data Processed Successfully")



	// error interface of builtin package
	// process_data := eprocess()

	// if process_data != nil{
	// 	fmt.Println("Error",process_data)
	// 	return
	// }
	// fmt.Println("Data Processed Successfully")

	err := readData()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully")

}
// using the error interface
type myError struct{
	message string

}

func (m *myError) Error() string{
	return fmt.Sprintf("Error:%s",m.message)
}

func eprocess() error{
	return &myError{"Custom error message"}
}

func readData() error{
	err := readConfig()
	if err != nil{
		return fmt.Errorf("readData:%w",err)
	}
	return nil
}
func readConfig() error{
	return errors.New("Config error")
}