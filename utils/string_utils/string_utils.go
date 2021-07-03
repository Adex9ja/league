package string_utils

import (
	"strings"
)

func JoinMatrix(matrix [][]string) []string {
	var strs []string
	for _, row := range matrix {
		s := strings.Join(row, ",")
		strs = append(strs, s)
	}
	return strs
}
