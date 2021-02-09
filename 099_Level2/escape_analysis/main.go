package main

func foo() int {
	x := 42
	return x
}

func bar() *int {
	y := 43
	return &y
}

// go run -gcflags "-m" main.go
func main() {
	_ = foo()
	_ = bar()
}
