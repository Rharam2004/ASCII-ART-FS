// package ascii 

// import (
// 	"fmt"
// )

// func filterEnglishCharacters(input string) error {
// 	filtered := ""
// 	for _, char := range input {
// 		if char >= 32 && char <= 126 {
// 			filtered += string(char)
// 		} else {
// 			return fmt.Errorf("Encountered non-English character: '%c'", char)
// 		}
// 	}
// 	return nil
// }

package ascii

import (
	"fmt"
	"regexp"
)

// Ascii checks if a given string contains non-ASCII characters.
func Ascii(str string) error {
	// Regular expression to match non-ASCII characters
	regex := regexp.MustCompile(`[^\x00-\x7F]`)
	if regex.MatchString(str) {
		return fmt.Errorf("error: the string contains non-ASCII characters")
	}
	return nil
}