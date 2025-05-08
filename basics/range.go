package basics

import "fmt"


func main(){
	message := "Hello World"

	for i,v := range message{
		fmt.Println(i,v)
		fmt.Printf("index:%d,Rune:%c\n",i,v)
	}
}