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
	CodeGetUserInfoFailed       = 406
	CodeLiveNotFound            = 407
	CodeLiveTitleEmpty          = 408
	CodeRecordStarted           = 409
	CodeUnknownError            = 400
)

var codeToMsg = map[int]string{
	CodeSuccess:                 "success",
	CodeUserExisted:             "user existed",
	CodeParamInvalid:            "param invalid",
	CodeUsernameOrPasswordError: "username or password error",
	CodeAuthCheckTokenFail:      "auth check token fail",
	CodeAuthTimeout:             "auth timeout",
	CodeLiveNotFound:            "live not found",
	CodeGetUserInfoFailed:       "get user info failed",
	CodeLiveTitleEmpty:          "live title empty",
	CodeRecordStarted:           "record already started",
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
