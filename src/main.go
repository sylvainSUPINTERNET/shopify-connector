package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	firebase "firebase.google.com/go"
	"context"
	types "github.com/sylvainSUPINTERNET/shopify-connector/types"
)


func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode) // use ReleaseMode


	r.POST("/webhook/cart/created", func (c *gin.Context) {
	

		opt := option.WithCredentialsFile("../accountCredentials.json")
		app, _ := firebase.NewApp(context.Background(), nil, opt)
		fcmClient, _ := app.Messaging(context.TODO())

		response, err := fcmClient.Send(context.Background(), &messaging.Message{
			Notification: &messaging.Notification{
				Title:    "New order - TEST",
				Body:     "Amount TEST",
				ImageURL: "https://www.wowisclassic.com/media/CACHE/images/wow/talents/21a8c903-6f70-49af-8a0e-1e8438b22956/a05b96f36dd18d6f90fabba0917da382.jpg",
			},
			// use one client web or mobile to generate this kind of token for fcm
			//Token: "dT8o52q92XObe5Y68c85oq:APA91bHwZPusjTqZz-qfUz32Ptgt-8Qe0xF9oO6WMg4RcHOkiEIO9GWODOgNbDUJ1GKVCse2DERJSFUe35qQYYNiWzpE-jIu3IVo9Z8nGy4TtxS5mbweeuUsB4M37yPYu7BU-HHUaXfn", // a token that you received from a client
			Token: "crWXZarvBBGsa72UYKtEgn:APA91bFqj5fJmxw4IvrwyQGb0J3YWLYkstwuTIWZVadkJjCCKzQL-GwcMYoflCeA36S6HnlMVho3cByrN8B8TaZa_zH3v_MHFzpVNXZYA9n_kwxSV0NXUQOBtIqklgmNPkuYer62xorz",
		})

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(response)
		}

		c.JSON(http.StatusOK, gin.H{"newOrder":"TEST"})
	});


	r.POST("/webhook/order", func (c *gin.Context) {
		var webhookOrder types.WebhookOrder
		err := c.BindJSON(&webhookOrder)
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
		}
		fmt.Println(webhookOrder.Email)

		opt := option.WithCredentialsFile("../accountCredentials.json")
		app, _ := firebase.NewApp(context.Background(), nil, opt)
		fcmClient, _ := app.Messaging(context.TODO())

		response, err := fcmClient.Send(context.Background(), &messaging.Message{
			Notification: &messaging.Notification{
				Title:    "New order - " + webhookOrder.Name,
				Body:     "Amount " + webhookOrder.TotalPrice,
				ImageURL: "https://www.wowisclassic.com/media/CACHE/images/wow/talents/21a8c903-6f70-49af-8a0e-1e8438b22956/a05b96f36dd18d6f90fabba0917da382.jpg",
			},
			// use one client web or mobile to generate this kind of token for fcm
			//Token: "dT8o52q92XObe5Y68c85oq:APA91bHwZPusjTqZz-qfUz32Ptgt-8Qe0xF9oO6WMg4RcHOkiEIO9GWODOgNbDUJ1GKVCse2DERJSFUe35qQYYNiWzpE-jIu3IVo9Z8nGy4TtxS5mbweeuUsB4M37yPYu7BU-HHUaXfn", // a token that you received from a client
			Token: "crWXZarvBBGsa72UYKtEgn:APA91bFqj5fJmxw4IvrwyQGb0J3YWLYkstwuTIWZVadkJjCCKzQL-GwcMYoflCeA36S6HnlMVho3cByrN8B8TaZa_zH3v_MHFzpVNXZYA9n_kwxSV0NXUQOBtIqklgmNPkuYer62xorz",
		})

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(response)
		}

		c.JSON(http.StatusOK, gin.H{"newOrder":webhookOrder})
	});


	r.GET("/test/notif", func ( c *gin.Context)  {
		opt := option.WithCredentialsFile("../accountCredentials.json")
		app, _ := firebase.NewApp(context.Background(), nil, opt)
		fcmClient, _ := app.Messaging(context.TODO())

		_, err := fcmClient.Send(context.Background(), &messaging.Message{
			Notification: &messaging.Notification{
				Title:    "A nice notification title",
				Body:     "A nice notification body",
			},
			// use one client web or mobile to generate this kind of token for fcm
			Token: "dT8o52q92XObe5Y68c85oq:APA91bHwZPusjTqZz-qfUz32Ptgt-8Qe0xF9oO6WMg4RcHOkiEIO9GWODOgNbDUJ1GKVCse2DERJSFUe35qQYYNiWzpE-jIu3IVo9Z8nGy4TtxS5mbweeuUsB4M37yPYu7BU-HHUaXfn", // a token that you received from a client
		})
		if err != nil {
			fmt.Println("fcm error")
		}
		c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	});


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// // https://github.com/firebase/snippets-go/blob/master/admin/main.go


// // => https://firebase.google.com/docs/cloud-messaging/send-message

// // https://stackoverflow.com/questions/47431757/fcm-http-v1-how-to-get-access-token-using-go

// import (
// 	_"reflect"
// 	_"log"
// 	// "fmt"
// 	"context"
// 	// firebase "firebase.google.com/go"
// 	// "firebase.google.com/go/messaging"
// 	// "google.golang.org/api/option"
// 	// _"github.com/gin-gonic/gin"
// 	// types "github.com/sylvainSUPINTERNET/shopify-connector/types"
// 	_"net/http"
// 	// _"github.com/davecgh/go-spew/spew"
// 	 "golang.org/x/oauth2/google"
// 	 "golang.org/x/oauth2"
// 	 "io/ioutil"
// 	 "errors"
// 	)


// func main() {

// 	// tokenProvider, err := newTokenProvider("../accountCredentials.json")
	
// 	// if err != nil {
// 	// 	fmt.Errorf(" Error firebase init : %v", err)
// 	// }
	

// 	// accessToken, _ := tokenProvider.token()
// 	// fmt.Println(accessToken);

// 	// opt := option.WithCredentialsFile("../accountCredentials.json")
// 	// app, _ := firebase.NewApp(context.Background(), nil, opt)

// 	// fcmClient, _ := app.Messaging(context.TODO())

// 	// response, err := fcmClient.Send(context.Background(), &messaging.Message{
// 	// 	Notification: &messaging.Notification{
// 	// 		Title:    "A nice notification title",
// 	// 		Body:     "A nice notification body",
// 	// 	},
// 	// 	Token: accessToken, // a token that you received from a client
// 	// })

// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// fmt.Println(response)

// 	// r := gin.Default()
// 	// gin.SetMode(gin.DebugMode) // use ReleaseMode

// 	// opt := option.WithCredentialsFile("../accountCredentials.json")
// 	// app, _ := firebase.NewApp(context.Background(), nil, opt)
	
// 	// fcmClient, _ := app.Messaging(context.TODO())
	
//     // spew.Dump(fcmClient)


// 	// if err != nil {
// 	// 	fmt.Errorf(" Error firebase init : %v", err)
// 	// 	//return nil, fmt.Errorf(" Error firebase init : %v", err)
// 	// }

// 	// fmt.Println("Firebase OK")

// 	// fcmClient, err := app.Messaging(context.TODO())

// 	// t := reflect.TypeOf(&fcmClient)
// 	// for i := 0; i < t.NumMethod(); i++ {
// 	// 	fmt.Println("FUCK")
// 	// 	m := t.Method(i)
// 	// 	fmt.Println(m.Name)
// 	// }

// 	// if err != nil {
// 	// 	log.Fatalf("messaging: %s", err) 
// 	// }

// 	// response, err := fcmClient.Send(context.Background(), &messaging.Message{
// 	// 	Notification: &messaging.Notification{
// 	// 		Title:    "A nice notification title",
// 	// 		Body:     "A nice notification body",
// 	// 	},
// 	// 	Token: "client-push-token", // a token that you received from a client
// 	// })
	
// 	// if err != nil {
// 	// 	fmt.Println("error fcm")
// 	// 	fmt.Println(err);
// 	// }
// 	// fmt.Println(response)

// 	// r.GET("/ping", func(c *gin.Context) {
// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"message": "pong",
// 	// 	})
// 	// })

// 	// /**
// 	// * Configured on shopify dashboard admin
// 	// **/
// 	// // TODO : check integrity shopify webhook
// 	// // TODO : send notif => https://stackoverflow.com/questions/37371990/how-can-i-send-a-firebase-cloud-messaging-notification-without-use-the-firebase
// 	// // TODO:  https://firebase.google.com/docs/cloud-messaging/send-message#go
// 	// // https://hackernoon.com/how-to-send-millions-of-push-notifications-with-go-and-firebase-cloud-messaging-554w35rs
// 	// r.POST("/webhook/order", func (c *gin.Context) {
// 	// 	var webhookOrder types.WebhookOrder
// 	// 	c.BindJSON(&webhookOrder)
// 	// 	fmt.Println(webhookOrder.Email)
// 	// 	c.JSON(http.StatusOK, gin.H{"newOrder":webhookOrder})
// 	// });
// 	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }