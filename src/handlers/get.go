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
	/*
		Init function runs before main
		Using aws-sdk creates a connection to Dynamodb
	*/
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("AWS connection failed: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session)
	}
}

func get(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/*
	   Handler function , used to retrive item from DynamoDB
	*/

	// get id value from the requests path parameters
	// get table name from environment variable set
	var (
		id        = request.PathParameters["id"]
		tableName = aws.String(os.Getenv("DEVICES_TABLE_NAME"))
	)

	// create input object dynamodb
	input := &dynamodb.GetItemInput{
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	// get item from dynamodb
	result, err := ddb.GetItem(input)

	// begin HTTP response
	if err != nil {
		body := "Internal Server Error"

		return events.APIGatewayProxyResponse{ // HTTP response internal server error
			Body:       string(body),
			StatusCode: 500,
		}, nil

	} else if result.Item == nil { // HTTP response record not found
		body := "Not Found"

		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 404,
		}, nil

	} else { // Success HTTP request
		device := Device{}
		dynamodbattribute.UnmarshalMap(result.Item, &device)
		device.Id = string(request.Path) + device.Id
		device.DeviceModel = string(request.Path) + device.DeviceModel
		body, _ := json.Marshal(device)

		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 200,
		}, nil

	}
}

func main() {
	lambda.Start(get)
}
