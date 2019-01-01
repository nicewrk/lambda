package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	OK      bool   `json:"ok"`
}

func Handler(r Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Processed request ID: %d", r.ID),
		OK:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
