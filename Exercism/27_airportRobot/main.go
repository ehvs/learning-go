package main

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(visitorName string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(visitorName))
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}

func (pt Portuguese) LanguageName() string {
	return "Portuguese"
}

func (pt Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

type German struct{}

func (gm German) LanguageName() string {
	return "German"
}

func (gm German) Greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}

func main() {
	fmt.Println(SayHello("Dietrich", German{}))
	fmt.Println(SayHello("Pierre", Italian{}))
	fmt.Println(SayHello("Felistino", Portuguese{}))
}
