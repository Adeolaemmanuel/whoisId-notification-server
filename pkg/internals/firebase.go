package internals

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

//Initiates firebase app
func initFirebaseApp() (*firebase.App) {
	const file string = "pkg/config/service_account_key.json"

	config := option.WithCredentialsFile(file)

	app, err := firebase.NewApp(context.Background(), nil, config)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

//Initiates messaging from firebase app
func Messaging() (*messaging.Client) {
	app := initFirebaseApp()

	messaging, err := app.Messaging(context.Background())
	
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return messaging
}