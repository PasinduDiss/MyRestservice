package main

import (
	handler "MyRestservice/src/handlers"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type App struct {
	handler handler.DeviceClient
}

func (a *App) GetHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	handler, error := a.handler.Get(ctx, request)

	return handler, error
}

func main() {
	var app App
	lambda.Start(app.GetHandler)
}
