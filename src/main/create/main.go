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

//CreateHandler function uses
func (a *App) CreateHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	handler, error := a.handler.Create(ctx, request)
	return handler, error
}

func main() {
	var app App
	lambda.Start(app.CreateHandler)
}
