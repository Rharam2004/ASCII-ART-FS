// package ascii

// import (

// 	"fmt"

// )

// func getBannerFilePath(banner string) string {
// 	var filePath string

// 	if banner == "standard" || banner == "" {
// 		filePath = "banners/standard.txt"
// 	} else if banner == "shadow" {
// 		filePath = "banners/shadow.txt"
// 	} else if banner == "thinkertoy" {
// 		filePath = "banners/thinkertoy.txt"
// 	} else {
// 		fmt.Println("Please enter correct banner name")
// 		return ""
// 	}

// 	return filePath
// }

package ascii

import (
	"fmt"
)

func GetBannerFilePath(banner string) string {
	var filePath string

	if banner == "standard" || banner == "" {
		filePath = "banners/standard.txt"
	} else if banner == "shadow" {
		filePath = "banners/shadow.txt"
	} else if banner == "thinkertoy" {
		filePath = "banners/thinkertoy.txt"
	} else {
		fmt.Println("Please enter correct banner name")
		return ""
	}

	return filePath
}
