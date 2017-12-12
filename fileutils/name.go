package fileutils

import (
	"fmt"
	"gocommon/stringutils"
	"strings"
)

func MakeTempFilePath(basedir string) (filepath string) {
	basedir = strings.Trim(basedir, " ")
	if strings.HasSuffix(basedir, "/") == true {
		filepath = fmt.Sprintf("%s%s", basedir, stringutils.Uuid4())
	} else {
		filepath = fmt.Sprintf("%s/%s", basedir, stringutils.Uuid4())
	}
	return filepath
}
