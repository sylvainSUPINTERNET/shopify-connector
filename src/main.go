package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	types "github.com/sylvainSUPINTERNET/shopify-connector/types"
	"net/http"
)

func main() {
	r := gin.Default()


	

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	/**
	* Configured on shopify dashboard admin
	**/
	// TODO : check integrity shopify webhook
	// TODO : send notif => https://stackoverflow.com/questions/37371990/how-can-i-send-a-firebase-cloud-messaging-notification-without-use-the-firebase
	// TODO:  https://firebase.google.com/docs/cloud-messaging/send-message#go
	// https://hackernoon.com/how-to-send-millions-of-push-notifications-with-go-and-firebase-cloud-messaging-554w35rs
	r.POST("/webhook/order", func (c *gin.Context) {
		var webhookOrder types.WebhookOrder
		c.BindJSON(&webhookOrder)
		fmt.Println(webhookOrder.Email)
		c.JSON(http.StatusOK, gin.H{"newOrder":webhookOrder})
	});
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}