package intermediate

import (
	"fmt"
	"regexp"
)


func main(){
fmt.Println("He said,\"I am great\"")
re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

email1 := "user@email.com"
email2 := "invalid_email"

// Match
fmt.Println("email1",re.MatchString(email1))
fmt.Println("email2",re.MatchString(email2))

// Capturing groups
// Compile a regex pattern to capture date components

re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

// test string
date := "2024-07-30"

// find all submatches

submatches := re.FindStringSubmatch(date)
fmt.Println(submatches)
fmt.Println(submatches[0])
fmt.Println(submatches[1])
fmt.Println(submatches[2])
fmt.Println(submatches[3])



}
