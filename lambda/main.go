package main

import (
	"lambda-func/app"
	"lambda-func/middleware"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func ProtectedHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "This is some top secret type sh",
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambdaApp := app.NewApp()
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return lambdaApp.ApiHandler.RegisterUser(request)
		case "/login":
			return lambdaApp.ApiHandler.LoginUser(request)
		case "/protected":
			// this is how to chain methods. the next keyword inside of the validate function is also involved here.
			return middleware.ValidateJWTMiddleware(ProtectedHandler)(request)
		default:
			return events.APIGatewayProxyResponse{
				Body:       "Not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}

	})
}
