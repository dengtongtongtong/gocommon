package webutils

type result struct {
	ErrorMessage string      `json:"errormesaage"`
	ErrorCode    int         `json:"errorcode"`
	Ret          interface{} `json:"ret"`
}

func Success(data interface{}) (ret result) {
	ret = result{ErrorMessage: "ok", Ret: data}
	return ret
}

func Failure(failure string, failurelist interface{}) (ret interface{}) {
	return ret
}
