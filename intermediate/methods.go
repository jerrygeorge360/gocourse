package intermediate

import "fmt"

type Rectangle struct{
	length float64
	width float64
}

// method with value receiver
func(r Rectangle) Area() float64{
	return r.length * r.width
}

func(r *Rectangle) ChangeLength(length float64){
	r.length = length
}

func main(){

	rect := Rectangle{length: 10, width: 8}
	rect.ChangeLength(2)
	area := rect.Area()

	fmt.Printf("%f\n",area)
	



}