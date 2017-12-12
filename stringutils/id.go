package stringutils

import (
	"github.com/satori/go.uuid"
)

func Uuid4() (ret string) {
	ret = uuid.NewV4().String()
	return ret
}
