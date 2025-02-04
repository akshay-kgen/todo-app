package repo

import (
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
func (r *TodoRepo) GetTodo(todo *models.TodoModel) (*models.TodoModel, error) {
	return nil, nil
}
