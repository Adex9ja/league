//Here the input from the controller is processed in this file
//and a response is sent back to the controller
package services

import (
	"encoding/csv"
	"fmt"
	"github.com/Adex9ja/solution/domain/matrix"
	"github.com/Adex9ja/solution/utils/arithemetic_utils"
	"github.com/Adex9ja/solution/utils/errors"
	"github.com/Adex9ja/solution/utils/string_utils"
	"strings"
)

var (
	MatrixService matrixServiceInterface = &matrixService{}
)

type matrixService struct {
}

type matrixServiceInterface interface {
	Echo(matrix matrix.Matrix) (string, *errors.ErrorResponser)
	Flatten(matrix matrix.Matrix) (string, *errors.ErrorResponser)
	Sum(matrix matrix.Matrix) (string, *errors.ErrorResponser)
	Multiply(matrix matrix.Matrix) (string, *errors.ErrorResponser)
	Invert(matrix matrix.Matrix) (string, *errors.ErrorResponser)
}

//This reads through a csv file and returns a response in a certain format
//Example
//Code:
//matrix := matrix.Matrix{fileData:file}
// string,err := services.MatrixService.Echo(matrix)
//Output
//1,2,3
//4,5,6
//7,8,9
func (m *matrixService) Echo(matrix matrix.Matrix) (string, *errors.ErrorResponser) {
	records, err := m.csvOpener(matrix)
	if err != nil {
		return "", err
	}
	var response string

	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	return response, nil
}

//This reads through a file and returns a response in a certain format
//Example
//Code:
//matrix := matrix.Matrix{fileData:file}
// string,err := services.MatrixService.Flatten(matrix)
//Output
//1,2,3,4,5,6,7,8,9
func (m *matrixService) Flatten(matrix matrix.Matrix) (string, *errors.ErrorResponser) {
	records, err := m.csvOpener(matrix)

	if err != nil {
		return "", err
	}

	strs := string_utils.JoinMatrix(records)
	response := fmt.Sprint(strings.Join(strs, ","))
	return response, nil
}

//This reads through a csv file and returns the sum of all inputs in the csv file
//Example
//Code:
//matrix := matrix.Matrix{fileData:file}
// string,err := services.MatrixService.Sum(matrix)
//Output
//int
func (m *matrixService) Sum(matrix matrix.Matrix) (string, *errors.ErrorResponser) {
	records, err := m.csvOpener(matrix)

	if err != nil {
		return "", err
	}

	var result string
	var sumErr *errors.ErrorResponser
	var appendedRow []string

	for _, k := range records {
		for _, w := range k {
			appendedRow = append(appendedRow, w)
		}
	}

	result, sumErr = arithemetic_utils.Sum(appendedRow)

	if sumErr != nil {
		return "", err
	}
	response := fmt.Sprint(result)
	return response, nil
}

//This reads through a csv file and returns the multiplication of all inputs in the csv file
//Example
//Code:
//matrix := matrix.Matrix{fileData:file}
// string,err := services.MatrixService.Multiply(matrix)
//Output
//int
func (m *matrixService) Multiply(matrix matrix.Matrix) (string, *errors.ErrorResponser) {
	records, err := m.csvOpener(matrix)

	if err != nil {
		return "", err
	}

	var result string
	var mulErr *errors.ErrorResponser
	var appendedRow []string

	for _, k := range records {
		for _, w := range k {
			appendedRow = append(appendedRow, w)
		}
	}

	result, mulErr = arithemetic_utils.Multiplication(appendedRow)

	if mulErr != nil {
		return "", err
	}
	response := fmt.Sprint(result)
	return response, nil
}

//This reads through a file and returns a response in a certain format
//Example
//Code:
//matrix := matrix.Matrix{fileData:file}
// string,err := services.MatrixService.Invert(matrix)
//Output
//string
func (m *matrixService) Invert(matrix matrix.Matrix) (string, *errors.ErrorResponser) {
	records, err := m.csvOpener(matrix)

	if err != nil {
		return "", err
	}

	var data [3][3]string
	var response string

	//Transposing the csv file
	for i, rows := range records {
		for j := range rows {
			data[j][i] = records[i][j]
		}
	}

	for _, row := range data {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row[:], ","))
	}

	return response, nil
}

//This reads the csv file inputted from the controller and returns a slice of slice
func (m *matrixService) csvOpener(matrix matrix.Matrix) ([][]string, *errors.ErrorResponser) {
	//Checks if the file being sent matches the accepted file and returns an error if it doesn't match
	if err := matrix.Validate(); err != nil {
		return [][]string{}, err
	}

	//Opens the file after the confirming the file matches the accepted files
	csvFile, err := matrix.FileData.Open()
	if err != nil {
		return [][]string{}, errors.NewNotFoundError("file not found")
	}

	//Reads through all the records in the file and returns [][]string
	records, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return [][]string{}, errors.NewInternalServerError("something went wrong")
	}

	//returns the data for being read and a nil as error
	return records, nil
}
