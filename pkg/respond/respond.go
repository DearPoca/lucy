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
	CodeGetUserInfoFailed       = 407
	CodeUnknownError            = 400
)

var codeToMsg = map[int]string{
	CodeSuccess:                 "success",
	CodeUserExisted:             "user existed",
	CodeParamInvalid:            "param invalid",
	CodeUsernameOrPasswordError: "username or password error",
	CodeAuthCheckTokenFail:      "auth check token fail",
	CodeAuthTimeout:             "auth timeout",
	CodeGetUserInfoFailed:       "get user info failed",
	CodeUnknownError:            "unknown error",
}

func CreateRespond(code int, data ...interface{}) *Respond {
	res := &Respond{Code: code, Msg: codeToMsg[code], Data: nil}
	if len(data) == 1 {
		res.Data = data[0]
	} else if len(data) > 1 {
		res.Data = data
	}
	return res
}
