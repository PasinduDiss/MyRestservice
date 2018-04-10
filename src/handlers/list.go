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

//Device struct used to represent a single device
type Device struct {
	ID          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}

//List Device Response
type ListDeviceResponse struct {
	Devices []Device `json:"devices"`
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

func list(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var (
		tableName = aws.String(os.Getenv("DEVICES_TABLE_NAME"))
	)

	input := &dynamodb.ScanInput{
		TableName: tableName,
	}
	result, _ := ddb.Scan(input)

	var devices []Device

	for _, i := range result.Items {
		device := Device{}
		if err := dynamodbattribute.UnmarshalMap(i, &device); err != nil {
			fmt.Println("Failed to unmarshal")
			fmt.Println(err)
		}
		devices = append(devices, device)
	}

	body, _ := json.Marshal(&ListDeviceResponse{
		Devices: devices,
	})
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(list)
}
