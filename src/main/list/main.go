package main

import (
	handler "MyRestservice/src/handlers"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//App struct
type App struct {
	Handler handler.Client
}

//ListHandler Returns lambda function from handler package
func (a *App) ListHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	handler, error := a.Handler.List(ctx, request)

	return handler, error
}

func main() {
	var devclient handler.DeviceClient
	app := &App{Handler: devclient}
	lambda.Start(app.ListHandler)
}
