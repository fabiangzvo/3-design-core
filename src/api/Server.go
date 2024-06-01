package api

import (
	"3-design-core/src/utils/logger"

	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
)

// InitServer initialize server
func InitServer() {
	const section = "server.InitServer"

	logger.Log.Infoln(section, "starting")

	router := gin.Default()

	router.Use(ginlogrus.Logger(logger.Log), gin.Recovery())

	Router(router)

	logger.Log.Infoln(section, "finished")
	router.Run(":8080")
}
