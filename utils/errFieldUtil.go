package utils

import (
	"net/http"

	"github.com/Gaurav-malwe/login-service/internal/constants"
)

var (
	LS1001 = setErrorFields(http.StatusUnauthorized, constants.AUTH_ERROR, "Auth Error")
)

func setErrorFields(httpStatus int, code string, errMessage string) map[string]interface{} {
	return map[string]interface{}{
		"HTTP_STATUS": httpStatus,
		"CODE":        code,
		"ERR_MESSAGE": errMessage,
	}
}

func CustomErrorFields(setErrorFields map[string]interface{}, customMessage string) map[string]interface{} {
	setErrorFields["ERR_MESSAGE"] = customMessage
	return setErrorFields
}
