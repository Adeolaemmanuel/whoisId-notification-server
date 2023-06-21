package notification

import "github.com/gin-gonic/gin"


func Notification(server *gin.RouterGroup) {
	server.POST("/notification/fcm", FcmHandler)
}