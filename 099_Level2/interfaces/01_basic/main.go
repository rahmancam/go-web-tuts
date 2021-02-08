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
	fmt.Printf("%T\n", a)
	a.speak()
}

func main() {
	dog := Dog{Name: "Woofie"}
	fmt.Printf("%T\n", dog)
	cat := Cat{Name: "Jessy"}
	fmt.Printf("%T\n", cat)
	talk(dog)
	talk(cat)
}
