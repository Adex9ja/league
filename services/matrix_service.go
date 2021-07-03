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

func (m *matrixService) Flatten(matrix matrix.Matrix) (string, *errors.ErrorResponser) {
	records, err := m.csvOpener(matrix)

	if err != nil {
		return "", err
	}

	strs := string_utils.JoinMatrix(records)
	response := fmt.Sprint(strings.Join(strs, ","))
	return response, nil
}

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

func (m *matrixService) Invert(matrix matrix.Matrix) (string, *errors.ErrorResponser) {
	records, err := m.csvOpener(matrix)

	if err != nil {
		return "", err
	}

	var data [3][3]string
	var response string

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

func (m *matrixService) csvOpener(matrix matrix.Matrix) ([][]string, *errors.ErrorResponser) {
	if err := matrix.Validate(); err != nil {
		return [][]string{}, err
	}

	csvFile, err := matrix.FileData.Open()
	if err != nil {
		return [][]string{}, errors.NewNotFoundError("file not found")
	}

	records, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return [][]string{}, errors.NewInternalServerError("something went wrong")
	}

	return records, nil
}
