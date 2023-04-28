package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound = "record not found"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func NoData(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

func BadRequest(c *gin.Context, message string, data ...interface{}) {
	obj := gin.H{"status": http.StatusBadRequest, "message": message}
	if len(data) > 0 {
		obj["data"] = data[0]
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, obj)
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Status:  http.StatusNotFound,
		Message: message,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": message})
}
