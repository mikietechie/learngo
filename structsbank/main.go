package main

import (
	"fmt"
	"strings"
)

const SYS_NAME = "Struct Bank"

type LivingThing interface {
	eat(food string)
}

type Person struct {
	name  string
	email string
}

func (person Person) Introduce() {
	line := fmt.Sprintf(`Hello world my name is %s`, person.name)
	fmt.Println(line)
}

func (person Person) eat(food string) {
	line := fmt.Sprintf(`I love eating %s`, food)
	fmt.Println(line)
}

func (person *Person) RespectMe() {
	person.name = strings.ToTitle(person.name)
	title := "SIR "
	if !strings.HasPrefix(person.name, title) {
		person.name = title + person.name
	}
}

func FeedIt(lt LivingThing) {
	lt.eat("Rice")
}

func main() {
	person := Person{name: "Mike", email: "mz@m.com"}
	person.RespectMe()
	person.RespectMe()
	person.Introduce()
	FeedIt(person)
}
