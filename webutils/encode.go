package webutils

import "encoding/base64"

func EncdoeURLWithBase64(data string) (ret string) {
	ret = base64.URLEncoding.EncodeToString([]byte(data))
	return ret
}

func DecodeURLWithBase64(data string) (ret []byte) {
	ret, _ = base64.URLEncoding.DecodeString(data)
	return ret
}
