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

	"github.com/satori/go.uuid"
)

type Devices struct {
	ID          string `json:"id"`
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
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session) // create dynamodb client
	}
}

func create(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Creating Device")

	var (
		id          = uuid.Must(uuid.NewV4(), nil).String()
		deviceModel = uuid.Must(uuid.NewV4(), nil).String()
		name        = uuid.Must(uuid.NewV4(), nil).String()
		note        = uuid.Must(uuid.NewV4(), nil).String()
		serial      = uuid.Must(uuid.NewV4(), nil).String()
		tableName   = aws.String(os.Getenv("DEVICES_TABLE_NAME"))
	)

	device := &Devices{
		ID:          id,
		DeviceModel: deviceModel,
		Name:        name,
		Note:        note,
		Serial:      serial,
	}

	//Parse requested body
	json.Unmarshal([]byte(request.Body), device)

	// Write to DynamoDB
	item, _ := dynamodbattribute.MarshalMap(device)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: tableName,
	}
	if _, err := ddb.PutItem(input); err != nil {
		return events.APIGatewayProxyResponse{ // Error HTTP response
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	} else {
		body, _ := json.Marshal(device)
		return events.APIGatewayProxyResponse{ // Success HTTP response
			Body:       string(body),
			StatusCode: 200,
		}, nil
	}
}

func main() {
	lambda.Start(create)
}
