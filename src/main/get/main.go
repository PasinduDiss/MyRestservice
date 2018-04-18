package main

import (
	handler "MyRestservice/src/handlers"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type App struct {
	Handler handler.Client
}

func (a *App) GetHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	handler, error := a.Handler.Get(ctx, request)

	return handler, error
}

func main() {
	var devclient handler.DeviceClient
	app := &App{Handler: devclient}
	lambda.Start(app.GetHandler)
}
