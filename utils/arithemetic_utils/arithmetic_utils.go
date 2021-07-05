//Mathematics operations are performed in this file
package arithemetic_utils

import (
	"github.com/Adex9ja/solution/utils/errors"
	"strconv"
)

//This sum the values in a given slice
//returns the sum and an error when it fails in the conversion of string to int
//Example
//Code
//row := []string{"1","2","3"}
//response,_ := arithemetic_utils.Sum(row)
//fmt.Println(response)
//Output
//6,nil
func Sum(row []string) (string, *errors.ErrorResponser) {
	sum := 0
	for _, s := range row {
		num, err := strconv.Atoi(s)

		if err != nil {
			return "", errors.NewBadRequestError("Not a number")
		}
		sum += num
	}
	return strconv.Itoa(sum), nil
}

//This performs the multiplication operation of a given slice
//returns the multiplication and an error when it fails in the conversion of string to int
//Example
//Code
//row := []string{"1","2","3"}
//response,_ := arithemetic_utils.Multiplication(row)
//fmt.Println(response)
//Output
//6,nil
func Multiplication(row []string) (string, *errors.ErrorResponser) {
	mul := 1
	for _, s := range row {
		num, err := strconv.Atoi(s)

		if err != nil {
			return "", errors.NewBadRequestError("Not a number")
		}
		mul *= num
	}

	return strconv.Itoa(mul), nil
}
