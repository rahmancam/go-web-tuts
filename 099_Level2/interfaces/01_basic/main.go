package main

import "fmt"

// Animal interface
type Animal interface {
	speak()
}

// Dog type
type Dog struct {
	Name string
}

// Cat Type
type Cat struct {
	Name string
}

func (d Dog) speak() {
	fmt.Println(d.Name + " => Woof.......woof")
}

func (c Cat) speak() {
	fmt.Println(c.Name + " => Meow.......meow")
}

func talk(a Animal) {
	a.speak()
}

func main() {
	dog := Dog{Name: "Woofie"}
	cat := Cat{Name: "Jessy"}
	talk(dog)
	talk(cat)
}
