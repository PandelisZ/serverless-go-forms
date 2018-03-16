package main

import (
	"fmt"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/schema"
	"github.com/pandelisz/serverless-go-forms/lib/helpers"
	"github.com/pandelisz/serverless-go-forms/lib/types"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	payload := types.ContactBasic{}
	urlencoded, err := url.ParseQuery(request.Body)
	decoder := schema.NewDecoder()

	if err != nil {
		fmt.Print(err)
		return events.APIGatewayProxyResponse{Body: helpers.ResponseFail("Could not decode url encoding"), StatusCode: 500}, nil
	}

	if err := decoder.Decode(&payload, urlencoded); err != nil {
		fmt.Print(err)

		return events.APIGatewayProxyResponse{Body: helpers.ResponseFail("Could not decode values"), StatusCode: 500}, nil
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
