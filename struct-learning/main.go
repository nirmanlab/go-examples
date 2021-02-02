package main

import (
	"fmt"
	"log"
)

func swap(x, y int) (int, int) {
	if x < 8 {
		log.Printf("%d", x)
	} else {
		fmt.Printf("%d", x)
	}
	return y, x
}

func main() {
	a, b := swap(7, 9)
	fmt.Println(a, b)
	// var reader interfacereader.Reader
	// reader = &random.FileReader{"File"}
	// reader.Read()
	// fmt.Printf("\n Asssign file reader %T \n", reader)
	// reader = &random.StreamReader{"File"}
	// reader.Read()
	// fmt.Printf("Assign stream reader %T \n", reader)
}
