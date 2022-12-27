package controller

type Respond struct {
	Code int
	Msg  string
}

var (
	ResSuccess     = &Respond{Code: 200, Msg: "success"}
	ResUserExisted = &Respond{Code: 200, Msg: "user existed"}

	ResUnknownError = &Respond{Code: 201, Msg: "unknown error"}
)
