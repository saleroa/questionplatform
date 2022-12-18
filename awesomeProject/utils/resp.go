package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResTemp struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var ok = ResTemp{
	Status: 200,
	Info:   "success",
}
var ParamErr = ResTemp{
	Status: 400,
	Info:   "param error",
}
var InterErr = ResTemp{
	Status: 500,
	Info:   "internal error",
}

func ResOk(c *gin.Context) {
	c.JSON(http.StatusOK, ok)
}

func ResParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamErr)
}

func ResInterErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InterErr)
}

func ResDesign(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
