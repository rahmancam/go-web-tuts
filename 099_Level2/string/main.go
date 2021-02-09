package main

import "fmt"

// Person type
type Person struct {
	Name string
	age  *int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %s", p.Name)
}

func main() {
	var age1 = 0
	var age2 = 0
	p1 := Person{"Abdul Rahman", &age1}
	fmt.Println(p1)

	p2 := Person{"Abdul Rahman", &age2}

	p3 := &p1

	fmt.Println(p1 == p2)
	fmt.Println(&p1 == &p2)
	fmt.Println(p2 == *p3)

}
