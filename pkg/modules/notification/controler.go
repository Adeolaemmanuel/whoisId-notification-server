package notification

import (
	"context"
	"whosId-notification/pkg/internals"

	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
)

type FcmPayload struct {
	Title string `json:"title"`
	Body string `json:"body"`
	Token []string `json:"token"`
}

// Handler called to handle fcm messaging
// for both single and multiple broadcast
func FcmHandler(ctx *gin.Context) {
	var payload FcmPayload

	if ctx.ShouldBindJSON(&payload) == nil {
		app := internals.Messaging()

		if len(payload.Token) == 1 {
			_, err :=app.Send(context.Background(), &messaging.Message{
				Notification: &messaging.Notification{
					Title: payload.Title,
					Body: payload.Body,
				},
				Token: payload.Token[0],
			})

			if err != nil {
				ctx.JSON(400, gin.H{
					"message": "Error sending push notification",
					"error": err,
				})
			}
		}

		if len(payload.Token) > 1 {
			_, err := app.SendMulticast(context.Background(), &messaging.MulticastMessage{
				Tokens: payload.Token,
				Notification: &messaging.Notification{
					Title: payload.Title,
					Body: payload.Body,
				},
			})

			if err != nil {
				ctx.JSON(400, gin.H{
					"message": "Error sending push notification",
					"error": err,
				})
			}
		}

		ctx.JSON(200, gin.H{
			"message": "Fcm notification sent successfully",
			"data": payload,
		})
	}
}