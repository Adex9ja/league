//String manipulation are performed in this file
package string_utils

import (
	"strings"
)

//Accepts a slice of slice as an argument
//Appends the slice of slice to return a slice
//Example

//Code:
// s := [][]string{[]string{"1","2","3"},[]string{"4","5","6"}}
// fmt.Print(string_utils.JoinMatrix(s))
//Output
// [1,2,3 4,5,6]
func JoinMatrix(matrix [][]string) []string {
	var strs []string
	for _, row := range matrix {
		s := strings.Join(row, ",")
		strs = append(strs, s)

	}
	return strs
}
