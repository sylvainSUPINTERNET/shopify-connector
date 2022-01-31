package main


// https://github.com/firebase/snippets-go/blob/master/admin/main.go


// => https://firebase.google.com/docs/cloud-messaging/send-message

// https://stackoverflow.com/questions/47431757/fcm-http-v1-how-to-get-access-token-using-go

import (
	_"reflect"
	_"log"
	"fmt"
	"context"
	//firebase "firebase.google.com/go"
	// _"firebase.google.com/go/messaging"
	// _"google.golang.org/api/option"
	// _"github.com/gin-gonic/gin"
	// types "github.com/sylvainSUPINTERNET/shopify-connector/types"
	_"net/http"
	// _"github.com/davecgh/go-spew/spew"
	 "golang.org/x/oauth2/google"
	 "golang.org/x/oauth2"
	 "io/ioutil"
	 "errors"
	)


const firebaseScope = "https://www.googleapis.com/auth/firebase.messaging"

type tokenProvider struct {
    tokenSource oauth2.TokenSource
}

// newTokenProvider function to get token for fcm-send
func newTokenProvider(credentialsLocation string) (*tokenProvider, error) {
    jsonKey, err := ioutil.ReadFile(credentialsLocation)
    if err != nil {
        return nil, errors.New("fcm: failed to read credentials file at: " + credentialsLocation)
    }
    cfg, err := google.JWTConfigFromJSON(jsonKey, firebaseScope)
    if err != nil {
        return nil, errors.New("fcm: failed to get JWT config for the firebase.messaging scope")
    }
    ts := cfg.TokenSource(context.Background())
    return &tokenProvider{
        tokenSource: ts,
    }, nil
}

func (src *tokenProvider) token() (string, error) {
    token, err := src.tokenSource.Token()
    if err != nil {
        return "", errors.New("fcm: failed to generate Bearer token")
    }
    return token.AccessToken, nil
}

func main() {

	tokenProvider, err := newTokenProvider("../accountCredentials.json")
	
	if err != nil {
		fmt.Errorf(" Error firebase init : %v", err)
	}
	

	accessToken, _ := tokenProvider.token()
	fmt.Println(accessToken);

	// r := gin.Default()
	// gin.SetMode(gin.DebugMode) // use ReleaseMode

	// opt := option.WithCredentialsFile("../accountCredentials.json")
	// app, _ := firebase.NewApp(context.Background(), nil, opt)
	
	// fcmClient, _ := app.Messaging(context.TODO())
	
    // spew.Dump(fcmClient)


	// if err != nil {
	// 	fmt.Errorf(" Error firebase init : %v", err)
	// 	//return nil, fmt.Errorf(" Error firebase init : %v", err)
	// }

	// fmt.Println("Firebase OK")

	// fcmClient, err := app.Messaging(context.TODO())

	// t := reflect.TypeOf(&fcmClient)
	// for i := 0; i < t.NumMethod(); i++ {
	// 	fmt.Println("FUCK")
	// 	m := t.Method(i)
	// 	fmt.Println(m.Name)
	// }

	// if err != nil {
	// 	log.Fatalf("messaging: %s", err) 
	// }

	// response, err := fcmClient.Send(context.Background(), &messaging.Message{
	// 	Notification: &messaging.Notification{
	// 		Title:    "A nice notification title",
	// 		Body:     "A nice notification body",
	// 	},
	// 	Token: "client-push-token", // a token that you received from a client
	// })
	
	// if err != nil {
	// 	fmt.Println("error fcm")
	// 	fmt.Println(err);
	// }
	// fmt.Println(response)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// /**
	// * Configured on shopify dashboard admin
	// **/
	// // TODO : check integrity shopify webhook
	// // TODO : send notif => https://stackoverflow.com/questions/37371990/how-can-i-send-a-firebase-cloud-messaging-notification-without-use-the-firebase
	// // TODO:  https://firebase.google.com/docs/cloud-messaging/send-message#go
	// // https://hackernoon.com/how-to-send-millions-of-push-notifications-with-go-and-firebase-cloud-messaging-554w35rs
	// r.POST("/webhook/order", func (c *gin.Context) {
	// 	var webhookOrder types.WebhookOrder
	// 	c.BindJSON(&webhookOrder)
	// 	fmt.Println(webhookOrder.Email)
	// 	c.JSON(http.StatusOK, gin.H{"newOrder":webhookOrder})
	// });
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}