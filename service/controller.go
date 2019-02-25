package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// CodeMessage ...
type CodeMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

// MonitorAddressList ...
func MonitorAddressList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		strings, e := store.LRange(RedisKeyNameIPFSSwarmAddress, 0, -1).Result()
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, strings)
	}
}

// MonitorAddressAdd ...
func MonitorAddressAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pins := ctx.PostFormArray("address")
		if pins != nil {
			var pdata []interface{}
			for _, v := range pins {
				pdata = append(pdata, v)
			}
			e := store.LPush(RedisKeyNameIPFSSwarmAddress, pdata...).Err()
			if e != nil {
				Error(ctx, e)
				return
			}
		}
		Success(ctx, pins)
	}
}

// MonitorPinsList ...
func MonitorPinsList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		strings, e := store.LRange(RedisKeyNameIPFSPins, 0, -1).Result()
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, strings)
	}
}

// MonitorPinsAdd ...
func MonitorPinsAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pins := ctx.PostFormArray("pins")
		if pins != nil {
			var pdata []interface{}
			for _, v := range pins {
				pdata = append(pdata, v)
			}
			e := store.LPush(RedisKeyNameIPFSPins, pdata...).Err()
			if e != nil {
				Error(ctx, e)
				return
			}
		}
		Success(ctx, pins)
	}
}

// Error ...
func Error(ctx *gin.Context, e error) {
	log.Error(e)
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
