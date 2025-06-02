package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	str := "Hello Go!"
	fmt.Println("Length:", len(str))

	// Slicing string (note: this could panic if slicing multi-byte characters)
	fmt.Println("Substring:", str[1:7])

	// String Conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println("Converted int to string:", str3)

	// Strings splitting
	fruits := "apple,orange,banana"
	parts := strings.Split(fruits, ",")
	fmt.Println("Split by comma:", parts)

	parts1 := strings.Split(fruits, "-")
	fmt.Println("Split by dash:", parts1)

	// Strings joining
	countries := []string{"Germany", "France", "UK"}
	joined := strings.Join(countries, ",")
	fmt.Println("Joined countries:", joined)

	// Contains substring
	contains := strings.Contains(str, "Go?")
	fmt.Println("Contains 'Go?':", contains)

	// Replace
	replaced := strings.Replace(str, "Go", "World", 1)
	fmt.Println("Replaced string:", replaced)

	// Remove leading/trailing spaces
	trimmed := strings.TrimSpace("   Hello Go!   ")
	fmt.Println("Trimmed:", trimmed)

	// Case conversion
	fmt.Println("Lowercase:", strings.ToLower(str))
	fmt.Println("Uppercase:", strings.ToUpper(str))

	// Repeat
	repeated := strings.Repeat("foo", 4)
	fmt.Println("Repeated:", repeated)

	// Count occurrences
	count := strings.Count("Hello", "l")
	fmt.Println("Count of 'l' in 'Hello':", count)

	// Prefix & Suffix
	fmt.Println("HasPrefix 'Hello':", strings.HasPrefix(str, "Hello"))
	fmt.Println("HasSuffix 'Go!':", strings.HasSuffix(str, "Go!"))

	// Regex: Find all numbers
	str5 := "Hello,123 Go!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println("Regex matched digits:", matches)

	// Using StringBuilder
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")
	result := builder.String()
	fmt.Println("StringBuilder result:", result)
}
