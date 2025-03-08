package main

import (
	"msantosfelipe/notification-receiver-lambda/config"
	handler "msantosfelipe/notification-receiver-lambda/handlers/lambda"
	"msantosfelipe/notification-receiver-lambda/infra"
	"msantosfelipe/notification-receiver-lambda/usecase"

	"github.com/aws/aws-lambda-go/lambda"
)

type Setup struct {
	notificationHandler handler.NotificationHandler
}

func setup() Setup {
	pushNotificationSender := infra.NewPushNotificationSender()
	uc := usecase.NewNotificationUsecase(pushNotificationSender)

	return Setup{
		notificationHandler: handler.NewNotificationHandler(uc),
	}
}

func main() {
	config.InitVars()
	setup := setup()

	lambda.Start(setup.notificationHandler)
}
