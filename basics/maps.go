package basics

import "fmt"

func main() {
	mymap := make(map[string]int)
	mymap["Key1"] = 9

	fmt.Println(mymap["Key1"])

}
