package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event map[string]interface{}) (string, error) {
	return fmt.Sprintf("Hello, %v!", event["name"]), nil
}

func main() {
	lambda.Start(Handler)
}
