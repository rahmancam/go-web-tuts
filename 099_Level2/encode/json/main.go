package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type person struct {
	Name string
}

func main() {
	p1 := person{
		Name: "Abdul Rahman",
	}

	p2 := person{
		Name: "Jamie",
	}

	persons := []person{p1, p2}

	bs, err := json.Marshal(persons)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(bs))

	persons2 := []person{}

	err = json.Unmarshal(bs, &persons2)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(persons2)

	json.NewEncoder(os.Stdout).Encode(persons)

	persons3 := []person{}

	json.NewDecoder(strings.NewReader(string(bs))).Decode(&persons3)
	fmt.Println(persons3)
}
