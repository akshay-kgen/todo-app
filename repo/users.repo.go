package repo

import (
	"fmt"

	"github.com/akshay-kgen/todo-app/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type UserRepo struct {
	Client *dynamodb.DynamoDB
}

func NewUserRepo(ddb *dynamodb.DynamoDB) *UserRepo {
	return &UserRepo{
		Client: ddb,
	}
}

func (r *UserRepo) CreateUser(user *models.UserModel) error {

	fmt.Println("userrrrr", user)
	data, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user to map: %v", err)
	}

	fmt.Println("dataaa", data)

	_, err = r.Client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("User"),
		Item:      data,
	})
	if err != nil {
		return fmt.Errorf("failed to create user in DynamoDB: %v", err)
	}

	return nil
}
