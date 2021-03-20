package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sachin-ghait-cld/bookstore_utils-go/logger"
)

var (
	router = gin.Default()
)

// StartApp starts the application
func StartApp() {
	mapUrls()
	logger.Info("Starting app on port 8090")
	router.Run(":8090")
}
