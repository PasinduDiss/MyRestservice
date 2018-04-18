package main

import (
	handler "MyRestservice/src/handlers"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

/*App struct is used by unit tests as well as the regular implementation of the
lambda functions*/
type App struct {
	Handler handler.Client
}

//CreateHandler function used to invoke the lambda fuction Create
func (a *App) CreateHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	handler, error := a.Handler.Create(ctx, request)
	return handler, error
}

func main() {
	var devclient handler.DeviceClient
	app := &App{Handler: devclient}
	lambda.Start(app.CreateHandler)
}
