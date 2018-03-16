package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pandelisz/serverless-go-forms/lib/helpers"
	"github.com/pandelisz/serverless-go-forms/lib/types"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	payload := types.ContactBasic{}

	if err := json.Unmarshal([]byte(request.Body), &payload); err != nil {
		fmt.Print(err)

		return events.APIGatewayProxyResponse{Body: helpers.ResponseFail(err.Error()), StatusCode: 500}, nil
	}
	payload.Referer = request.Headers["referer"]

	forwarder := helpers.Forwarder{
		Mailgun: true,
		Slack:   true,
	}

	if err := forwarder.Send(payload); err != nil {
		fmt.Print(err)

		return events.APIGatewayProxyResponse{Body: helpers.ResponseFail("An error occured while trying to send your message."), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: helpers.ResponseSuccess(), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
