package main

import (
	"log"

	"github.com/nickalie/go-webpbin"
)

func main() {
	err := webpbin.NewCWebP().
		Quality(60).
		InputFile("demo.jpg").
		OutputFile("image.webp").
		Run()
	if err != nil {
		log.Fatal(err)
	}

}
