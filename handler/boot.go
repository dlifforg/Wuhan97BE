package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"hello/service"
	"net/http"
)

type ErrorResponse struct {
	Error string
}
type Response struct {
	Offset int32
	Limit int32
	Prefix string
	Data interface{}
}

func Index(c *gin.Context)  {
	var req = struct {
		Offset int32  `form:"offset"`
		Limit  int32  `form:"limit"   binding:"required"`
		Prefix string `form:"prefix"  binding:"required"`
	}{}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	op := redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: int64(req.Offset),
		Count:  int64(req.Limit),
	}
	lists, err := service.Redis().ZRevRangeByScore(req.Prefix, op).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	var data []interface{}
	for _, i := range lists {
		var d interface{}
		json.Unmarshal([]byte(i),&d)
		data=append(data,d)
	}
	rsp:=Response{
		Offset: req.Offset,
		Limit:  req.Limit,
		Prefix: req.Prefix,
		Data:   data,
	}
	c.JSON(http.StatusOK, rsp)
}
