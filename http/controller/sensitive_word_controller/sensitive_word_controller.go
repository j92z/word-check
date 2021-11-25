package sensitive_word_controller

import (
	"github.com/gin-gonic/gin"
	"sensitive_words_check/http/http_util"
	"sensitive_words_check/service/sensitive_word_service"
)

type operateWordParam struct {
	Word string `json:"word" binding:"required,min=1"`
}

// AddSensitiveWord
// @Summary 添加词
// @Description 添加词
// @Tags Word
// @Accept  json
// @Produce  json
// @Param info body operateWordParam true "add word"
// @Success 200 {object} http_util.HttpResponse
// @Router /sensitive_word_service [put]
func AddSensitiveWord(c *gin.Context) {
	var param operateWordParam
	if err := c.ShouldBindJSON(&param); err != nil {
		http_util.ResponseFail(c, err)
		return
	}
	if err := sensitive_word_service.AddSensitiveWord(param.Word); err != nil {
		http_util.ResponseFail(c, err)
		return
	}
	http_util.ResponseOk(c, nil)
}

// RemoveSensitiveWord
// @Summary 移除词
// @Description 移除词
// @Tags Word
// @Accept  json
// @Produce  json
// @Param info body operateWordParam true "delete word"
// @Success 200 {object} http_util.HttpResponse
// @Router /sensitive_word_service [delete]
func RemoveSensitiveWord(c *gin.Context) {
	var param operateWordParam
	if err := c.ShouldBindJSON(&param); err != nil {
		http_util.ResponseFail(c, err)
		return
	}
	if err := sensitive_word_service.RemoveSensitiveWord(param.Word); err != nil {
		http_util.ResponseFail(c, err)
		return
	}
	http_util.ResponseOk(c, nil)
}

type checkWordParam struct {
	Words []string `json:"words" binding:"required,gt=0"`
}
// CheckSensitiveWord
// @Summary 检查词是否敏感
// @Description 检查词是否敏感
// @Tags Word
// @Produce  json
// @Param x-ldap-user header string true "ldap userid 仅在开发环境需要提供"
// @Param info query checkWordParam true "check words"
// @Success 200 {object} http_util.HttpResponse{info=sensitive_word_service.CheckSensitiveWordResult}
// @Router /sensitive_word_service [get]
func CheckSensitiveWord(c *gin.Context) {
	var param checkWordParam
	if err := c.ShouldBindJSON(&param); err != nil {
		http_util.ResponseFail(c, err)
		return
	}
	result := sensitive_word_service.CheckSensitiveWord(param.Words)
	http_util.ResponseOk(c, result)
}
