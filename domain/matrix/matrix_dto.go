package matrix

import (
	"github.com/Adex9ja/solution/utils/errors"
	"github.com/Adex9ja/solution/utils/file_utils"
	"mime/multipart"
)

type Matrix struct {
	FileData *multipart.FileHeader `form:"file"`
}

func (m *Matrix) Validate() *errors.ErrorResponser {

	extension := file_utils.GetFileType(m.FileData.Filename)

	if extension != ".csv" {
		return errors.NewBadRequestError("Please enter a csv file")
	}
	return nil
}
