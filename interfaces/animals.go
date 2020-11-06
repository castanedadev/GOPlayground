package main

import "fmt"

// Animal interface along with its methods
type Animal interface {
	move() string
	makeSound() string
}

// Define animal structs
type dog struct{}
type duck struct{}
type cat struct{}

// Implement common methods in interface depending on struct type
func (dog) move() string {
	return ("Je me promène comme un chien dans le parc")
}

func (duck) move() string {
	return ("Je me promène comme un canard dans le parc")
}

func (cat) move() string {
	return ("Je me promène comme un chat dans le parc")
}

func (dog) makeSound() string {
	return ("Woof! Woof!")
}

func (duck) makeSound() string {
	return ("Cuack! Cuack!")
}

func (cat) makeSound() string {
	return ("Miau! Miau!")
}

func move(a Animal) {
	fmt.Println(a.move())
}

func makeSound(a Animal) {
	fmt.Println(a.makeSound())
}

func main() {
	tuti := dog{}
	felix := cat{}
	psyduck := duck{}
	move(tuti)
	makeSound(psyduck)
	move(felix)
	makeSound(tuti)
}
