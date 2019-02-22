package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// CodeMessage ...
type CodeMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

// MonitorAddress ...
func MonitorAddress(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		strings, e := store.LRange(RedisKeyNameIPFSSwarmAddress, 0, -1).Result()
		if e != nil {
			logrus.Error(e)
			Error(ctx, e)
			return
		}
		Success(ctx, strings)
	}
}

// MonitorPins ...
func MonitorPins(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		strings, e := store.LRange(RedisKeyNameIPFSPins, 0, -1).Result()
		if e != nil {
			logrus.Error(e)
			Error(ctx, e)
			return
		}
		Success(ctx, strings)
	}
}

// Success ...
func Error(ctx *gin.Context, e error) {
	ctx.JSON(http.StatusOK, &CodeMessage{
		Code:    -1,
		Message: e.Error(),
		Detail:  nil,
	})
}

// Success ...
func Success(ctx *gin.Context, detail []string) {
	ctx.JSON(http.StatusOK, &CodeMessage{
		Code:    0,
		Message: "success",
		Detail:  detail,
	})
}
