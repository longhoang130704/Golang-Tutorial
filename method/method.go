package main

import (
	"fmt"
)

type receiver string

type Person struct {
	name string
	age int
}

func (receiver receiver) String() string {
	return "Receiver is non-struct"
}

func (p *Person) String() string {
	return fmt.Sprintf("Receiver is struct Person: %s is %d years old", p.name, p.age)
}

func main() {
	person := &Person{"Long", 20}
	information := person.String()
	fmt.Println(information)
	var nonStructReceiver receiver = "non-struct"
	result := nonStructReceiver.String()
	fmt.Println(result)
	fmt.Println("Study method") 
}
