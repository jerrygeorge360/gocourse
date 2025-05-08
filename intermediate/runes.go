package intermediate

import (
	"fmt"
	"unicode/utf8"
)


func main(){
	message := "Hello World \r go"

	rawMessage := `Hello World \n go`

	fmt.Println(message)
	fmt.Println(rawMessage)
	fmt.Println(len(message))
	fmt.Println(len(rawMessage))
	fmt.Println(message[0])

	for i,char := range message{
		fmt.Printf("Character at index %d is %c\n",i,char)
	}

	fmt.Println(utf8.RuneCountInString(message))

	var ch rune = 'a'

}