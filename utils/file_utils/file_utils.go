//File manipulations are performed in this file
package file_utils

import (
	"path/filepath"
)

// This takes a fileName as an argument
// returns the extension of the given file
//Example
//Code:
// fileName := "document.csv"
// fmt.Println(file_utils.GetFileType(fileName))
//Output
// ".csv"
func GetFileType(f string) string {
	extension := filepath.Ext(f)
	return extension
}
