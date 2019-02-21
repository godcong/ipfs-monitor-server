package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CodeMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

func MonitorAddress(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var detail []string

		Success(ctx, detail)
	}
}

func Success(ctx *gin.Context, detail []string) {
	ctx.JSON(http.StatusOK, &CodeMessage{
		Code:    0,
		Message: "success",
		Detail:  detail,
	})
}
