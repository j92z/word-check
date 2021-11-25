package http_util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sensitive_words_check/constant/operate_const"
)

type HttpResponse struct {
	Code    operate_const.OperateStatusType `json:"code" enums:"0,1"`
	Message string                          `json:"message"`
	Info    interface{}                     `json:"info"`
}

func Response(c *gin.Context, httpCode int, status operate_const.OperateStatusType, err error, info interface{}) {
	message := "success"
	if err != nil {
		message = err.Error()
	}
	c.JSON(httpCode, HttpResponse{
		Code:    status,
		Message: message,
		Info:    info,
	})
}

func ResponseOk(c *gin.Context, info interface{}) {
	Response(c, http.StatusOK, operate_const.Success, nil, info)
}

func ResponseFail(c *gin.Context, err error) {
	Response(c, http.StatusOK, operate_const.Fail, err, nil)
}
