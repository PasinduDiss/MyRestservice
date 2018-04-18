package main

import (
	handler "ServerlessRestAPI/src/handlers"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//App struct
type App struct {
	handler handler.DeviceClient
}

//ListHandler Returns lambda function from handler package
func (a *App) ListHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	handler, error := a.handler.List(ctx, request)

	return handler, error
}

func main() {
	var app App
	lambda.Start(app.ListHandler)
}
