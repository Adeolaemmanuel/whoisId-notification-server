package api

import (
	"whosId-notification/pkg/modules/notification"

	"github.com/gin-gonic/gin"
)

// Routes for all v1 api request
func Routes() {
	var server = gin.New()
	v1 := server.Group("/api/v1")

	notification.Notification(v1)

	server.Run(":4000")
}