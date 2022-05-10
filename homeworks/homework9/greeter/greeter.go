package greeter

import (
	"fmt"
	"strings"
)

func Greet(name string, hour int) string {
	greeting := "Hello"

	if hour >= 6 && hour < 12 {
		greeting = "Good morning"
	} else if hour >= 12 && hour < 18 {
		greeting = "Hello"
	} else if hour >= 18 && hour < 24 {
		greeting = "Good evening"
	} else if hour >= 0 && hour < 6 {
		greeting = "Good night"
	}
	trimmedName := strings.Trim(name, " ")
	return fmt.Sprintf("%s %s!", greeting, strings.Title(trimmedName))
}
