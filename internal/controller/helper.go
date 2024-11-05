package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Gaurav-malwe/login-service/internal/constants"
	"github.com/Gaurav-malwe/login-service/internal/model"
	"github.com/Gaurav-malwe/login-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func checkError(ginCtx *gin.Context, err error) bool {
	if err != nil {
		writeErrorOnResponse(ginCtx.Writer, utils.CustomErrorFields(utils.LS1001, err.Error()))
		return true
	}
	return false
}

func writeErrorOnResponse(responseWriter http.ResponseWriter, fields map[string]interface{}) {
	httpStatus, _ := fields["HTTP_STATUS"].(int)
	additionalMessage, _ := fields["ADDITIONAL_MESSAGE"].(string)

	response := model.StandardError{
		Version:        constants.VERSION,
		HttpStatusCode: httpStatus,
		Errors: []model.APIErrorResponse{
			{
				Code:              fields["CODE"].(string),
				Message:           fields["ERR_MESSAGE"].(string),
				AdditionalMessage: additionalMessage,
			},
		},
	}

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(httpStatus)
	json.NewEncoder(responseWriter).Encode(response)
}

func listErrors(err error) string {
	var arr []string
	for _, err := range err.(validator.ValidationErrors) {
		arr = append(arr, err.Field())
	}
	str := strings.Join(arr, ", ")
	return str
}
