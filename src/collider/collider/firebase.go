package collider

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

func newFirebase() *firebase.App {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://mona-44092-default-rtdb.europe-west1.firebasedatabase.app",
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("/home/niels/dev/mona/apprtc/src/collider/service_account.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("restricted_access/secret_document")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)

	return app
}

