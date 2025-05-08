package basics

import "fmt"

func main() {
	process()

}
func ultimate(){
	fmt.Println("hello")
}
func process() {
	defer ultimate()
	defer fmt.Println("Deferreed statement executed")
	defer fmt.Println("second Deferreed statement executed")
	defer fmt.Println("third Deferreed statement executed")
	fmt.Println("Normal execution")
}
