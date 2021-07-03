package arithemetic_utils

import (
	"fmt"
	"github.com/Adex9ja/solution/utils/errors"
	"strconv"
)

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

func Multiplication(row []string) (string, *errors.ErrorResponser) {
	sum := 1
	for _, s := range row {
		num, err := strconv.Atoi(s)

		if err != nil {
			return "", errors.NewBadRequestError("Not a number")
		}
		sum *= num
	}
	fmt.Println(sum)
	return strconv.Itoa(sum), nil
}
