package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString("hi.. there")
	file.Close()
	stream, err := os.ReadFile("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	readString := string(stream)
	fmt.Println(readString)

}
