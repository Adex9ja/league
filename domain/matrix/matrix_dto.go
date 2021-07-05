//This is the core of the app
//The entity to be used by the app is defined here
package matrix

import (
	"github.com/Adex9ja/solution/utils/errors"
	"github.com/Adex9ja/solution/utils/file_utils"
	"mime/multipart"
)

//The entity
type Matrix struct {
	FileData *multipart.FileHeader `form:"file"`
}

/*
Validates the type of file being inputted from the controller
returns an error when file doesn't match the accepted file
*/
func (m *Matrix) Validate() *errors.ErrorResponser {

	extension := file_utils.GetFileType(m.FileData.Filename)

	if extension != ".csv" {
		return errors.NewBadRequestError("Please enter a csv file")
	}
	return nil
}
