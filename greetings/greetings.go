package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello " + name)
	return message
}