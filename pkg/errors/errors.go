package errors

import "errors"

var (
	ErrStreamInactive             = errors.New("stream inactive")
	ErrStreamFormatError          = errors.New("stream format error")
	ErrLiveNotFound               = errors.New("live not found")
	ErrUserHaveNoLive             = errors.New("user have no live")
	ErrLiveFormatError            = errors.New("live name format error")
	ErrRequesterNotOwner          = errors.New("requester are not owner")
	ErrRecordingStarted           = errors.New("recording has been started")
	ErrTelephoneNumNotValid       = errors.New("telephone number not valid")
	ErrEmailFormatError           = errors.New("email format error")
	ErrUsernameOrPasswordNotValid = errors.New("username or password format not valid")
)
