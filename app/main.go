package main

import (
	"fmt"
	"msantosfelipe/notification-receiver-lambda/config"
	handler "msantosfelipe/notification-receiver-lambda/handlers/lambda"
	"msantosfelipe/notification-receiver-lambda/infra"
	"msantosfelipe/notification-receiver-lambda/usecase"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Setup struct {
	notificationHandler handler.NotificationHandler
}

func main() {
	config.InitVars()
	setup := setup()

	if config.ENV.IS_LOCAL {
		fmt.Println("Running in local mode...")
		runLocal(setup)
	} else {
		lambda.Start(setup.notificationHandler.ProcessNotification)
	}
}

func setup() Setup {
	pushNotificationSender := infra.NewPushNotificationSender()
	uc := usecase.NewNotificationUsecase(pushNotificationSender)

	return Setup{
		notificationHandler: handler.NewNotificationHandler(uc),
	}
}

func runLocal(setup Setup) {
	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{handler.ApiKeyValidationHeader: config.ENV.VALID_API_KEY},
		Body:    "{\"title\":\"Testetitle\",\"body\":\"Testebody\",\"app\":\"XP Investimentos\"}",
	}
	response, err := setup.notificationHandler.ProcessNotification(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
