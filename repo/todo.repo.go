package repo

import (
	"errors"
	"fmt"

	"github.com/akshay-kgen/todo-app/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type TodoRepo struct {
	Client    *dynamodb.DynamoDB
	GSI       string
	TableName string
}

var ErrTodoNotFound = errors.New("todo not found")

func NewTodoRepo(ddb *dynamodb.DynamoDB) *TodoRepo {
	return &TodoRepo{
		Client:    ddb,
		GSI:       "UserIndex",
		TableName: "Todo",
	}
}

func (r *TodoRepo) CreateTodo(todo *models.TodoModel) error {
	av, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item:      av,
	}

	_, err = r.Client.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepo) GetAllTodo(userId string) ([]*models.TodoModel, error) {

	input := &dynamodb.QueryInput{
		TableName:              aws.String("Todo"),
		IndexName:              aws.String(r.GSI),
		KeyConditionExpression: aws.String("userId = :userId"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userId": {
				S: aws.String(userId),
			},
		},
	}

	result, err := r.Client.Query(input)
	if err != nil {
		return nil, fmt.Errorf("failed to query todos: %w", err)
	}

	var todos []*models.TodoModel
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &todos)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal query result: %w", err)
	}

	return todos, nil
}

func (r *TodoRepo) GetTodo(userId, todoId string) (*models.TodoModel, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Todo"),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(userId),
			},
			"todoId": {
				S: aws.String(todoId),
			},
		},
	}

	result, err := r.Client.GetItem(input)
	if err != nil {
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	if result.Item == nil {
		return nil, ErrTodoNotFound
	}

	var todo models.TodoModel
	err = dynamodbattribute.UnmarshalMap(result.Item, &todo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal result: %w", err)
	}

	return &todo, nil
}
