package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var ddb *dynamodb.DynamoDB

func init() {
	/*
		Init function runs before main
		Using aws-sdk creates a connection to Dynamodb
	*/
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("AWS connection failed: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session) // Create DynamoDB client
	}
}

func Delete(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	/*
	 Implemets the Delete lambda function linked to the API's DELETE request
	 deleting one item from the dynamodb
	*/

	var (
		id        = request.PathParameters["id"]
		tableName = aws.String(os.Getenv("DEVICES_TABLE_NAME"))
	)

	// input created for dynamodb.DeleteItem
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: tableName,
	}

	//Delete item frm dynamodb
	_, err := ddb.DeleteItem(input)

	if err != nil {

		return events.APIGatewayProxyResponse{ //HTTP response internal server error
			Body:       err.Error(),
			StatusCode: 500,
		}, nil

	} else {
		return events.APIGatewayProxyResponse{ //HTTP Success response
			StatusCode: 203,
		}, nil
	}

}

func main() {
	lambda.Start(Delete)
}
