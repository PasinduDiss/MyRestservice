// src/handlers/create.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Device struct {
	Id          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}

var ddb *dynamodb.DynamoDB

func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("AWS connection failed: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session)
	}
}

func Create(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var (
		tableName = aws.String(os.Getenv("DEVICES_TABLE_NAME"))
	)

	// Initialize device struct
	device := &Device{
		Id:          "",
		DeviceModel: "",
		Name:        "",
		Note:        "",
		Serial:      "",
	}

	// Parse request body
	json.Unmarshal([]byte(request.Body), device)

	// Write to DynamoDB
	item, _ := dynamodbattribute.MarshalMap(device)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: tableName,
	}

	if _, err := ddb.PutItem(input); err != nil { // HTTP response internal server error

		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, nil

	} else { // HTTP Success response, Item created in dynamodb
		return events.APIGatewayProxyResponse{
			StatusCode: 201,
		}, nil
	}
}

func main() {
	lambda.Start(Create)
}
