package respond

type Respond struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResSuccess(data ...interface{}) *Respond {
	dataSlice := []interface{}(data)
	res := &Respond{Code: 200, Msg: "success", Data: nil}
	if len(dataSlice) == 1 {
		res.Data = dataSlice[0]
	} else if len(dataSlice) > 1 {
		res.Data = dataSlice
	}
	return res
}

func ResUserExisted(data ...interface{}) *Respond {
	dataSlice := []interface{}(data)
	res := &Respond{Code: 200, Msg: "user existed", Data: nil}
	if len(dataSlice) == 1 {
		res.Data = dataSlice[0]
	} else if len(dataSlice) > 1 {
		res.Data = dataSlice
	}
	return res
}

func ResParamInvalid(data ...interface{}) *Respond {
	dataSlice := []interface{}(data)
	res := &Respond{Code: 200, Msg: "param invalid", Data: nil}
	if len(dataSlice) == 1 {
		res.Data = dataSlice[0]
	} else if len(dataSlice) > 1 {
		res.Data = dataSlice
	}
	return res
}

func ResUsernameOrPasswordError(data ...interface{}) *Respond {
	dataSlice := []interface{}(data)
	res := &Respond{Code: 200, Msg: "username or password error", Data: nil}
	if len(dataSlice) == 1 {
		res.Data = dataSlice[0]
	} else if len(dataSlice) > 1 {
		res.Data = dataSlice
	}
	return res
}

func ResAuthCheckTokenFail(data ...interface{}) *Respond {
	dataSlice := []interface{}(data)
	res := &Respond{Code: 200, Msg: "auth check token fail", Data: nil}
	if len(dataSlice) == 1 {
		res.Data = dataSlice[0]
	} else if len(dataSlice) > 1 {
		res.Data = dataSlice
	}
	return res
}

func ResAuthTimeout(data ...interface{}) *Respond {
	dataSlice := []interface{}(data)
	res := &Respond{Code: 200, Msg: "auth timeout", Data: nil}
	if len(dataSlice) == 1 {
		res.Data = dataSlice[0]
	} else if len(dataSlice) > 1 {
		res.Data = dataSlice
	}
	return res
}

func ResBucketNameInvalidError(data ...interface{}) *Respond {
	dataSlice := []interface{}(data)
	res := &Respond{Code: 200, Msg: "bucket name invalid", Data: nil}
	if len(dataSlice) == 1 {
		res.Data = dataSlice[0]
	} else if len(dataSlice) > 1 {
		res.Data = dataSlice
	}
	return res
}

func ResUnknownError(data ...interface{}) *Respond {
	dataSlice := []interface{}(data)
	res := &Respond{Code: 200, Msg: "unknown error", Data: nil}
	if len(dataSlice) == 1 {
		res.Data = dataSlice[0]
	} else if len(dataSlice) > 1 {
		res.Data = dataSlice
	}
	return res
}
