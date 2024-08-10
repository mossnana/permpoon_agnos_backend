package response

import (
	"github.com/gin-gonic/gin"
)

func HandleResponse(c *gin.Context, status int, e error) {
	c.JSON(status, gin.H{
		"error": e.Error(),
	})
}
