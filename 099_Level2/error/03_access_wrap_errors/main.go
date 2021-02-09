package main

import (
	"errors"
	"fmt"
	"net"
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

func openFile(filename string) (string, error) {
	return "", ErrFileNotFound{
		Filename: filename,
		Base:     ErrNotExist,
	}
}

func processFile(filename string) error {
	_, err := openFile(filename)

	if err != nil {
		return fmt.Errorf("Processing file '%s' failed: %w", filename, err)
	}

	return nil
}

func main() {
	err := processFile("secret.txt")
	if err != nil {
		if errors.Is(err, ErrNotExist) {
			fmt.Println("This is an ErrNotExist")
		}
		var fErr ErrFileNotFound
		if errors.As(err, &fErr) {
			fmt.Printf("Was unable to do something with file %s\n", fErr.Filename)
		}
		var netErr net.Error
		if errors.As(err, &netErr) {
			if netErr.Temporary() {
				// Retry request
			} else {
				// some other error
			}
		}
		fmt.Println(err)
	}
}
