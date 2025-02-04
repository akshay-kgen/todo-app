package main

import (
	"fmt"
	"net/http"

	"github.com/akshay-kgen/todo-app/config"
	"github.com/akshay-kgen/todo-app/routes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type App struct {
	router http.Handler
	ddb    *dynamodb.DynamoDB
	config *config.Config
}

func NewApp(configI *config.Config) *App {

	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   &configI.AwsConfig.Region,
			Endpoint: &configI.DynamoEndpoint,
			Credentials: credentials.NewStaticCredentials(
				configI.AwsConfig.AccessKey,
				configI.AwsConfig.SecretKey,
				"",
			),
		},
	}))

	app := &App{
		ddb:    dynamodb.New(awsSession),
		config: configI,
	}

	app.loadRoutes()

	return app
}

func (app *App) Start() error {

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.PORT),
		Handler: app.router,
	}

	_, err := app.ddb.ListTables(&dynamodb.ListTablesInput{Limit: aws.Int64(1)})
	if err != nil {
		return fmt.Errorf("error connecting db: %w", err)
	}

	// create table script
	// scripts.CreateDynamodbTables(app.config)

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error running server: %w", err)
	}

	return nil
}

func (app *App) loadRoutes() {
	app.router = routes.NewRoutes(app.ddb)
}
