package helper

import (
	"net/http"
	"tourism-monitoring/constant"
)

func GetResponseCodeFromErr(err error) int {
	switch err {
	case constant.EMAIL_NOT_FOUND:
		return http.StatusBadRequest
	case constant.EMAIL_IS_EMPTY:
		return http.StatusBadRequest
	case constant.PASSWORD_IS_EMPTY:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
