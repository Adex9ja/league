package file_utils

import (
	"path/filepath"
)

func GetFileType(f string) string {
	extension := filepath.Ext(f)
	return extension
}
