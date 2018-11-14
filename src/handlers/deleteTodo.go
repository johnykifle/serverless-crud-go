package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

var ddb *dynamodb.DynamoDB

func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session)
	}
}

// DeleteTodo : delete a todo item
func DeleteTodo(ctx context.Context, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	fmt.Println("DeleteTodo")

	var (
		id        = request.PathParameters["id"]
		tableName = aws.String(os.Getenv("TODOS_TABLE_NAME"))
	)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: tableName,
	}
	_, err := ddb.DeleteItem(input)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode: 204,
		}
	}
}

func main() {
	lambda.Start(DeleteTodo)
}
