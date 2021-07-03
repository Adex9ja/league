package controllers

import (
	"github.com/Adex9ja/solution/domain/matrix"
	"github.com/Adex9ja/solution/services"
	"github.com/Adex9ja/solution/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Echo(c *gin.Context) {

	var mtx matrix.Matrix

	file, err := c.FormFile("file")
	if err != nil {
		errorBody := errors.NewBadRequestError("no file is received")
		c.JSON(errorBody.Status, errorBody)
		return
	}

	mtx.FileData = file

	result, saveErr := services.MatrixService.Echo(mtx)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.String(http.StatusOK, result)
}

func Flatten(c *gin.Context) {
	var mtx matrix.Matrix

	file, err := c.FormFile("file")
	if err != nil {
		errorBody := errors.NewBadRequestError("no file is received")
		c.JSON(errorBody.Status, errorBody)
		return
	}

	mtx.FileData = file

	result, saveErr := services.MatrixService.Flatten(mtx)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.String(http.StatusOK, result)
}

func Sum(c *gin.Context) {
	var mtx matrix.Matrix

	file, err := c.FormFile("file")
	if err != nil {
		errorBody := errors.NewBadRequestError("no file is received")
		c.JSON(errorBody.Status, errorBody)
		return
	}

	mtx.FileData = file

	result, saveErr := services.MatrixService.Sum(mtx)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.String(http.StatusOK, result)
}

func Multiply(c *gin.Context) {
	var mtx matrix.Matrix

	file, err := c.FormFile("file")
	if err != nil {
		errorBody := errors.NewBadRequestError("no file is received")
		c.JSON(errorBody.Status, errorBody)
		return
	}

	mtx.FileData = file

	result, saveErr := services.MatrixService.Multiply(mtx)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.String(http.StatusOK, result)
}

func Invert(c *gin.Context) {
	var mtx matrix.Matrix

	file, err := c.FormFile("file")
	if err != nil {
		errorBody := errors.NewBadRequestError("no file is received")
		c.JSON(errorBody.Status, errorBody)
		return
	}

	mtx.FileData = file

	result, saveErr := services.MatrixService.Invert(mtx)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.String(http.StatusOK, result)
}
