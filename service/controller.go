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
		log.Info("address list:", strings)
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, strings)
	}
}

// MonitorAddressDelete ...
func MonitorAddressDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		address := ctx.PostForm("address")
		c, e := store.LRem(RedisKeyNameIPFSSwarmAddress, 0, address).Result()
		log.Info("address del:", c)
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, c)
	}
}

// MonitorAddressAdd ...
func MonitorAddressAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		address := ctx.PostFormArray("address")
		log.Info("address add:", address)
		if address != nil {
			var pdata []interface{}
			for _, v := range address {
				pdata = append(pdata, v)
			}
			e := store.LPush(RedisKeyNameIPFSSwarmAddress, pdata...).Err()
			if e != nil {
				Error(ctx, e)
				return
			}
		}
		Success(ctx, address)
	}
}

// MonitorPinsList ...
func MonitorPinsList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		strings, e := store.LRange(RedisKeyNameIPFSPins, 0, -1).Result()
		log.Info("pins list:", strings)
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, strings)
	}
}

// MonitorAddressDelete ...
func MonitorPinsDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pins := ctx.PostForm("pins")
		c, e := store.LRem(RedisKeyNameIPFSPins, 0, pins).Result()
		log.Info("pins del:", c)
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, c)
	}
}

// MonitorPinsAdd ...
func MonitorPinsAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pins := ctx.PostFormArray("pins")
		log.Info("pins add:", pins)
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
func Success(ctx *gin.Context, detail interface{}) {
	ctx.JSON(http.StatusOK, &CodeMessage{
		Code:    0,
		Message: "success",
		Detail:  detail,
	})
}
