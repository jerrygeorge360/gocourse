package intermediate

import "fmt"


func main(){
	// printing actions
	// fmt.Print("Hello")
	// fmt.Print(222)

	// fmt.Println("Hello")
	// fmt.Println(344)

	// name := "Jerry"
	// age := 28
	// fmt.Printf("Name:%s,Age:%d\n",name,age)

	// scanning functions
	var name string
	var age int

	fmt.Print("Enter your name and age:")
	// fmt.Scan(&name,&age)
	// fmt.Scanln(&name)
	// fmt.Scanln(&age)
	fmt.Scanf("%s,%d",&name,&age)
	fmt.Printf("Name:%s Age:%d",name,age)

	// Error Formatting functions
	err := checkAge(15)
	if err != nil{
		fmt.Println("Error ",err)
	}	

	
}

func checkAge(age int) error{
	if age<18{
		return fmt.Errorf("Age %d is too young to drive",age)
	}
	return nil
}