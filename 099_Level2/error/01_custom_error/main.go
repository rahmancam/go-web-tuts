package main

import (
	"errors"
	"fmt"
)

// ErrFileNotFound error
type ErrFileNotFound struct {
	Filename string
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("File %s was not found", e.Filename)
}

// Is checks for type ErrFileNotFound
func (e ErrFileNotFound) Is(other error) bool {
	_, ok := other.(ErrFileNotFound)
	return ok
}

// ErrNotExist file not exist
var ErrNotExist = fmt.Errorf("File does not exist")

// ErrUserNotExist user not exist
var ErrUserNotExist = errors.New("User does not exist")

func main() {
	// err := ErrNotExist
	// err := ErrUserNotExist
	err := ErrFileNotFound{
		Filename: "secret.txt",
	}

	fmt.Println(errors.Is(err, ErrFileNotFound{}))

	if err == ErrNotExist {
		fmt.Println("The requested File not found")
	} else if err == ErrUserNotExist {
		fmt.Println("You need to register first")
	} else {
		fmt.Println(err)
	}
}
