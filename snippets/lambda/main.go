// GOOS=linux go build -o main
package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message string  `json:"message"`
	Status  float64 `json:"statusCode"`
}

func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprint("Process request id", request.ID),
		Status:  200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
