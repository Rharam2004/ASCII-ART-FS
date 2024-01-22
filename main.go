package main

import (
	ascii "ascii/form"

	"fmt"

	"os"

	"strings"
)

func main() {
	if len(os.Args) <= 1 || len(os.Args) > 3 {
		fmt.Println("Error: Incorrect number of arguments")
		return
	}

	arg := os.Args[1]
	banner := "standard"
	if len(os.Args) > 3 {
		banner = os.Args[2]
	}

	if arg == "" {

		return

	}
	// Call the Ascii function from the 'ascii' package to validate the input string
	err := ascii.Ascii(arg)
	if err != nil {
		fmt.Println(err)
		return
	}
	newlines := strings.Split(arg, "\\n")

	if ascii.Check(newlines) {
		for range newlines[:(len(newlines)-1)]{
			fmt.Println()
		}
		return
	}
	line := ""
	for _, word := range newlines {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					line += ascii.Lines(2 + int(char-' ')*9 + i, banner)
				}
				if line != "" {
					fmt.Println(line)
					line = ""
				}
			}
		}
		if word == "" {
			fmt.Println()
		}
	}
}
