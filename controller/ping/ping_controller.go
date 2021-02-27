package ping

import (
	"github.com/gin-gonic/gin"
)

// Ping allows ping
func Ping(c *gin.Context) {
	// c.String(http.StatusOK, "pong")
	c.JSON(200, gin.H{"msg": "yup", "app": "alive"})
}
