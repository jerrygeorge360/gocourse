package intermediate

import "fmt"

type person struct{
	name string
	age int
}

type Employee struct{
	person // Embedded struct
	empId string
	salary float64
}

func main(){
	emp := Employee{
		person: person{name: "John",age: 30},
		empId: "E021",
		salary: 190000,
	}
	fmt.Println("Name",emp.name)

}