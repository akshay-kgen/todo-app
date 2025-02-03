package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/akshay-kgen/todo-app/routes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type App struct {
	router http.Handler
	ddb    *dynamodb.DynamoDB
}

func NewApp() *App {
	region := os.Getenv("AWS_REGION")
	ddbEndpoint := os.Getenv("DYNAMO_ENDPOINT")

	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   &region,
			Endpoint: &ddbEndpoint,
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_ACCESS_KEY"),
				os.Getenv("AWS_SECRET_KEY"),
				"",
			),
		},
	}))

	app := &App{
		ddb: dynamodb.New(awsSession),
	}

	app.loadRoutes()

	return app
}

func (app *App) Start() error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: app.router,
	}

	_, err := app.ddb.ListTables(&dynamodb.ListTablesInput{Limit: aws.Int64(1)})
	if err != nil {
		return fmt.Errorf("error connecting db: %w", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error running server: %w", err)
	}

	return nil
}

func (app *App) loadRoutes() {
	app.router = routes.NewRoutes(app.ddb)
}
