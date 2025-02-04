package scripts

import (
	"fmt"

	"github.com/akshay-kgen/todo-app/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateDynamodbTables(configI *config.Config) {
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

	ddb := dynamodb.New(awsSession)

	createUserTable(ddb)

	createTodoTable(ddb)

}

func createUserTable(ddb *dynamodb.DynamoDB) {
	input := &dynamodb.CreateTableInput{
		TableName: aws.String("User"),

		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("userId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("email"),
				AttributeType: aws.String("S"),
			},
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("userId"),
				KeyType:       aws.String("HASH"),
			},
		},

		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("UserEmailIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("email"),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
		},

		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := ddb.CreateTable(input)
	if err != nil {
		fmt.Println("Error creating user table:", err)
		return
	}

	fmt.Println("TodoApp table created successfully!")
}

func createTodoTable(ddb *dynamodb.DynamoDB) {
	input := &dynamodb.CreateTableInput{
		TableName: aws.String("Todo"),

		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("todoId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("userId"),
				AttributeType: aws.String("S"),
			},
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("todoId"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("userId"),
				KeyType:       aws.String("RANGE"),
			},
		},

		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("UserIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("userId"),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := ddb.CreateTable(input)
	if err != nil {
		fmt.Println("Error creating todo table:", err)
		return
	}

	fmt.Println("Todo table created successfully!")
}
