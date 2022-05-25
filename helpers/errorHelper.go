package helpers

import (
	"github.com/gin-gonic/gin"
)

type errorStruct struct {
	Error string
}

func MyAbort(c *gin.Context, str string) {
	c.AbortWithStatusJSON(400, errorStruct{Error: str})
}

var (
	otpChars = "1234567890"
)
