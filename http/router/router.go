package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "sensitive_words_check/docs"
	"sensitive_words_check/http/controller/sensitive_word_controller"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	//健康检查
	r.GET("/ping", ping)
	sensitiveWordApi := r.Group("sensitive_word_service")
	{
		sensitiveWordApi.POST("", sensitive_word_controller.AddSensitiveWord)
		sensitiveWordApi.DELETE("", sensitive_word_controller.RemoveSensitiveWord)
		sensitiveWordApi.GET("", sensitive_word_controller.CheckSensitiveWord)
	}
	r.GET("/api-doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}


// 健康检查
// @Summary PING TEST
// @Description test service useful
// @Tags Common
// @Success 200 {string} pong
// @Router /ping [get]
func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}