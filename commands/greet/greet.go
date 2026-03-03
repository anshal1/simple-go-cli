package greet

import (
	"fmt"
	"math/rand"
	"strings"
)

var randomGreetings = [5]string{"Hello", "Good Morning", "Good Night", "Nice to meet you", "Happy to see you"}

func joinValues(args map[string]string) string {
	var values strings.Builder
	for key, value := range args {
		if key == "command" || key == "flags" {
			continue
		}
		values.WriteString(" " + value)
	}
	return values.String()
}

func Greet(args map[string]string) {
	var greetString strings.Builder
	greetString.WriteString("Hello")
	greetString.WriteString(joinValues(args))
	fmt.Println(greetString.String())
}

func GreetWithFlags(args map[string]string) {
	flags := args["flags"]
	var response strings.Builder
	response.WriteString("Hello" + joinValues(args))
	for _, flag := range flags {
		if flag == '-' {
			continue
		}
		if flag == 'r' {
			response.Reset()
			randomNumber := rand.Intn(len(randomGreetings))
			response.WriteString(randomGreetings[randomNumber])
			response.WriteString(joinValues(args))
		}
		if flag == 'c' {
			var tempval = response.String()
			response.Reset()
			response.WriteString(strings.ToUpper(tempval))
		}
	}
	fmt.Println(response.String())
}
