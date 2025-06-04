package intermediate

import (
	"os"
	"text/template"
)


func main(){

	templ := template.New("example")
	templ,err := templ.Parse("Welcome, {{.name}}! How are you doing?")
	if err != nil{
		panic(err)
	}

	data := map[string]interface{}{
		"name":"John",
	}
	err = templ.Execute(os.Stdout,data)
	if err != nil{
		panic(err)
	}
}