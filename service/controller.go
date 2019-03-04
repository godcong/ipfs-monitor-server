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

		strings, e := store.HGetAll(RedisKeyNameIPFSSwarmAddress).Result()
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
		i, e := store.HDel(RedisKeyNameIPFSSwarmAddress, address).Result()
		log.Info("address del:", i)
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, i)
	}
}

// MonitorAddressAdd ...
func MonitorAddressAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		address := ctx.PostFormArray("address")
		log.Info("address add:", address)
		if address != nil {
			for _, v := range address {
				e := store.HSet(RedisKeyNameIPFSSwarmAddress, v, 0).Err()
				if e != nil {
					Error(ctx, e)
					return
				}
			}
		}
		Success(ctx, address)
	}
}

// MonitorPinsList ...
func MonitorPinsList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		strings, e := store.HGetAll(RedisKeyNameIPFSPins).Result()
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
		i, e := store.HDel(RedisKeyNameIPFSPins, pins).Result()
		log.Info("pins del:", i)
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, i)
	}
}

// MonitorPinsAdd ...
func MonitorPinsAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pins := ctx.PostFormArray("pins")
		log.Info("pins add:", pins)
		if pins != nil {
			for _, v := range pins {
				e := store.HSet(RedisKeyNameIPFSPins, v, 0).Err()
				if e != nil {
					Error(ctx, e)
					return
				}
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
