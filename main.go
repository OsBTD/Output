package main

import (
	"fmt"
	"os"

	"ascii/ascii"
)

func main() {
	args := os.Args[1:]
	// we get the banner style and input text from user input
	filename, input, Banner, inputsplit := ascii.FileNameManagement()
	// we read the ASCII art characters from the chosen banner file
	ascii.ReadText(Banner)

	// we populate the map with ASCII characters
	replaceMap := ascii.Populate()

	// and finally we print the resulting ascii art
	ascii.PrintArt(input, inputsplit, replaceMap, filename)
	fmt.Println("File processed successfully :) ")

	fmt.Println("filename is :", filename, "\ninput is :", input, "\nBanner is :", Banner, "\nargs len is :", len(args))
}
