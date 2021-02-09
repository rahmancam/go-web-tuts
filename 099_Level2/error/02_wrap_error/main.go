package main

import (
	"errors"
	"fmt"
)

// ErrFileNotFound error
type ErrFileNotFound struct {
	Filename string
	Base     error
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("File %s was not found, %v", e.Filename, e.Base)
}

func (e ErrFileNotFound) Unwrap() error {
	return e.Base
}

// ErrNotExist file not exist
var ErrNotExist = fmt.Errorf("File does not exist")

// ErrUserNotExist user not exist
var ErrUserNotExist = errors.New("User does not exist")

func openFile(filename string) (string, error) {
	return "", ErrNotExist
}

func openFile2(filename string) (string, error) {
	return "", ErrFileNotFound{
		Filename: filename,
		Base:     ErrNotExist,
	}
}

func main() {
	_, err := openFile("test.txt")
	if err != nil {
		wrappedErr := fmt.Errorf("Unable to open file %v: %w", "test.txt", err)
		if errors.Is(wrappedErr, ErrNotExist) {
			fmt.Println("This is an ErrNotExist")
		}
		fmt.Println(wrappedErr)
	}

	_, err = openFile2("secret.txt")
	if err != nil {
		if errors.Is(err, ErrNotExist) {
			fmt.Println("This is an ErrNotExist")
		}
		fmt.Println(err)
	}
}
