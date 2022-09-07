package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello Documentation!")
}

func ExampleNew() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}
