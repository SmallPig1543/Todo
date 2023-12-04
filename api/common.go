package api

import (
	"Todo/pkg/ctl"
	"Todo/pkg/e"
	"encoding/json"
)

func ErrorResponse(err error) *ctl.Response {
	_, ok := err.(*json.UnmarshalTypeError)
	if ok {
		return ctl.RespError(err, e.InvalidParams)
	}
	return ctl.RespError(err, e.ERROR)
}
