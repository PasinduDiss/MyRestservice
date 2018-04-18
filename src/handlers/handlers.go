package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//Device struct is used to represent the body recieved
type Device struct {
	Id          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}

//ListDeviceResponse struct is used to represent the body recieved
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

//Client interface created for the mainfunction to access lambda functions
type Client interface {
	Get(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Create(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Delete(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	List(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

// DeviceClient is a structure
type DeviceClient struct{}

//Create lambda function POSTs items to dynamodb
func (d DeviceClient) Create(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

//ValidateInput helps validate the input revieved by the Create lambda function
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

//Delete lambda function DELETEs a specified item from dynamodb
func (d DeviceClient) Delete(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

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

//Get lambda function GETs a specified item to dynamodb
func (d DeviceClient) Get(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/*
		   Implemets the create lambda function linked to the API's GET request
			 returning one item from the dynamodb
	*/

	// get id value from the request's path parameters
	// get table name from environment variables
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
		device.Id = string(request.Path)
		device.DeviceModel = "/deviceModel/" + device.DeviceModel
		body, _ := json.Marshal(device)

		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 200,
		}, nil

	}
}

//List lambda function retrives all items from dynamodb
func (d DeviceClient) List(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/*
		Implemets the list lambda function linked to the API's GET request,
		which returns the list of items in the dynamodb table
	*/
	var (
		tableName = aws.String(os.Getenv("DEVICES_TABLE_NAME"))
	)
	//Creating necesary input fo dynamodb.Scan function
	input := &dynamodb.ScanInput{
		TableName: tableName,
	}
	//Get contents of dynamodb table
	result, _ := ddb.Scan(input)

	var devices []Device

	for _, i := range result.Items {
		/*loop through items recived by API Gateway, create Device objects
		appened to deviced list
		*/
		device := Device{}
		if err := dynamodbattribute.UnmarshalMap(i, &device); err != nil {
			fmt.Println("Failed to unmarshal")
			fmt.Println(err)
		}
		devices = append(devices, device)
	}

	body, _ := json.Marshal(&ListDeviceResponse{ // convert devices list in to json
		Devices: devices,
	})
	return events.APIGatewayProxyResponse{ // HTTP Success
		Body:       string(body),
		StatusCode: 200,
	}, nil
}
