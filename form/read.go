package ascii

import (
	"bufio"

	"fmt"

	"os"

	"strings"
)

func Lines(integers int, banner string) string {
    // Prepare an empty string to store the line of text
    lines := ""
	var file *os.File
    var errors error
    
    file, errors = os.Open("banners/" + banner + ".txt")
    if errors != nil {
        fmt.Print(errors)
        os.Exit(0)
    }

    // Create a reader tool to read the contents of a file
    reader := bufio.NewReader(file)
    
    // Read lines from the file until we reach the desired line number
    for r := 0; r < integers; r++ {
        lines, _ = reader.ReadString('\n')
    }
    
    // Remove the trailing newline character, if any
    lines = strings.TrimSuffix(lines, "\n")
    
    // Return the line of text
    return lines
}