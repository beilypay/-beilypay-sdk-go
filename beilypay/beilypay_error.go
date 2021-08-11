package beilypay

import "strconv"

type BeilypayError struct {
	ErrorCode int //错误码 500失败
	ErrorMsg  string
}

func (be *BeilypayError) Error() string {
	return "ErrorCode:" + strconv.Itoa(be.ErrorCode) + "," + "ErrorMsg:" + be.ErrorMsg
}
