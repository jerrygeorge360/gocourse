package main
import "fmt"

func main(){

	process()
	fmt.Println("Returned from Process")

}

func process(){
	defer func(){
		if r:= recover();r != nil{
			fmt.Println("Recovered",r)
		}
	}()
	fmt.Println("start process")
	panic("somethihg went wrong")
	fmt.Println("end process")
}