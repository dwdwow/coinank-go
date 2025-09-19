package coinank

import "errors"

const (
	// 1 is success
	ErrCodeSystem                       = 0
	ErrCodeTooFrequent                  = -2
	ErrCodeInvalidApikey                = -3
	ErrCodeNoSubscription               = -4
	ErrCodeNoAccessRightOfThisTimeLevel = -5
	ErrCodeNoAccessRightOfThisAPI       = -6
	ErrCodeOutOfTimeRange               = -7
)

var (
	ErrSystem                       = errors.New("system error")
	ErrTooFrequent                  = errors.New("too frequent")
	ErrInvalidApikey                = errors.New("invalid apikey")
	ErrNoSubscription               = errors.New("no subscription")
	ErrNoAccessRightOfThisTimeLevel = errors.New("no access right of this time level")
	ErrNoAccessRightOfThisAPI       = errors.New("no access right of this api")
	ErrOutOfTimeRange               = errors.New("out of time range")
)

func ErrFromCode(code int) error {
	switch code {
	case ErrCodeSystem:
		return ErrSystem
	case ErrCodeTooFrequent:
		return ErrTooFrequent
	case ErrCodeInvalidApikey:
		return ErrInvalidApikey
	case ErrCodeNoSubscription:
		return ErrNoSubscription
	case ErrCodeNoAccessRightOfThisTimeLevel:
		return ErrNoAccessRightOfThisTimeLevel
	case ErrCodeNoAccessRightOfThisAPI:
		return ErrNoAccessRightOfThisAPI
	case ErrCodeOutOfTimeRange:
		return ErrOutOfTimeRange
	}
	return errors.New("unknown error")
}
