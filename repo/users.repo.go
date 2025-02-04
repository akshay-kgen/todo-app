package repo

import (
	"errors"
	"fmt"

	"github.com/akshay-kgen/todo-app/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type UserRepo struct {
	Client *dynamodb.DynamoDB
	GSI    string
}

func NewUserRepo(ddb *dynamodb.DynamoDB) *UserRepo {
	return &UserRepo{
		Client: ddb,
		GSI:    "UserEmailIndex",
	}
}

func (r *UserRepo) CreateUser(user *models.UserModel) error {

	data, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user to map: %v", err)
	}

	_, err = r.Client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("User"),
		Item:      data,
	})
	if err != nil {
		return fmt.Errorf("failed to create user in DynamoDB: %v", err)
	}

	return nil
}

func (r *UserRepo) GetUserByEmail(email string) (*models.UserModel, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String("User"),
		IndexName:              aws.String(r.GSI),
		KeyConditionExpression: aws.String("email = :email"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {S: aws.String(email)},
		},
		Limit: aws.Int64(1),
	}

	result, err := r.Client.Query(input)
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, errors.New("user not found")
	}

	var user models.UserModel
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
