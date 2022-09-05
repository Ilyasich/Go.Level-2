package main

import (
	"fmt"
	"os"
)

func main() {
	emptyFile, err := os.Create("File")
	for i := 0; i < 1000000; i++ {
		fmt.Println(emptyFile)
		if err != nil {
			panic(err)
		}
		defer func(emptyFile *os.File) {
			if e := recover(); e != nil {
				fmt.Println("no file: ", e)
				err = e.(error)
			}
			err := emptyFile.Close()
			if err != nil {

			}
		}(emptyFile)
		{
			if err := recover(); err != nil {
				fmt.Println("error", err)
			}

			fmt.Println(emptyFile, "new file")
		}

	}
}
