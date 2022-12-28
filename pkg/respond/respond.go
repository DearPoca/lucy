package respond

type Respond struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

const (
	CodeSuccess                 = 200
	CodeUserExisted             = 401
	CodeParamInvalid            = 402
	CodeUsernameOrPasswordError = 403
	CodeAuthCheckTokenFail      = 404
	CodeAuthTimeout             = 405
	CodeBucketNameInvalidError  = 406
	CodeUnknownError            = 400
)

var codeToMsg = map[int]string{
	CodeSuccess:                 "success",
	CodeUserExisted:             "user existed",
	CodeParamInvalid:            "param invalid",
	CodeUsernameOrPasswordError: "username or password error",
	CodeAuthCheckTokenFail:      "auth check token fail",
	CodeAuthTimeout:             "auth timeout",
	CodeBucketNameInvalidError:  "bucket name invalid",
	CodeUnknownError:            "unknown error",
}

func CreateRespond(code int, data ...interface{}) *Respond {
	dataSlice := []interface{}(data)
	res := &Respond{Code: code, Msg: codeToMsg[code], Data: nil}
	if len(dataSlice) == 1 {
		res.Data = dataSlice[0]
	} else if len(dataSlice) > 1 {
		res.Data = dataSlice
	}
	return res
}
