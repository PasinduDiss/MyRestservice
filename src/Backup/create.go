// src/handlers/create.go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"

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

func Create(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/*
	   implemets the create lambda function linked to the API's POST request
	*/

	var (
		tableName = aws.String(os.Getenv("DEVICES_TABLE_NAME"))
	)

	// Initialize device struct
	device := &Device{}

	// Parse request body
	json.Unmarshal([]byte(request.Body), device)

	// Validate data retrieved from request body
	validationerr := ValidateInput(device)

	if validationerr != nil {
		// HTTP Bad request, if not all fields are included
		return events.APIGatewayProxyResponse{
			Body:       "Bad request",
			StatusCode: 400,
		}, nil
	}

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

func ValidateInput(device *Device) error {
	/*
		   Validate device feilds are not empty or missing
			 Validate the input to dynamodb removes the resource path in
			 Id and DeviceModel Attributes
	*/
	var ErrMissingField = errors.New("Missing Attribute field")

	if device.Id == "" {
		return ErrMissingField
	}
	if device.DeviceModel == "" {
		return ErrMissingField
	}
	if device.Name == "" {
		return ErrMissingField
	}
	if device.Note == "" {
		return ErrMissingField
	}
	if device.Serial == "" {
		return ErrMissingField
	}
	var re = regexp.MustCompile(`^\/.*\/`)
	device.Id = re.ReplaceAllString(device.Id, "")
	device.DeviceModel = re.ReplaceAllString(device.DeviceModel, "")
	return nil
}

func main() {
	lambda.Start(Create)
}
