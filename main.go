package main

// Import necessary packages.
import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// MyEvent represents the structure of the incoming Lambda event.
type MyEvent struct {
	Name string `json:"what is your name?"` // Field for the name in the event.
	Age  int    `json:"How old are you?"`   // Field for the age in the event.
}

// MyResponse represents the structure of the response returned by the Lambda function.
type MyResponse struct {
	Message string `json:"Answer:"` // Field for the response message.
}

// HandleLambdaEvent is the main handler function for the Lambda event.
// It takes a MyEvent as input and returns a MyResponse.
func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	// Format a response message based on the input event.
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", event.Name, event.Age)}, nil
}

// The main function is the entry point for the Lambda function.
func main() {
	// Start the Lambda function with the specified event handler.
	lambda.Start(HandleLambdaEvent)
}
