package infra

import (
	"fmt"
	"msantosfelipe/notification-receiver-lambda/config"

	"github.com/gregdel/pushover"
)

type pushNotificationSender struct {
	pushOverClient *pushover.Pushover
}

type PushNotificationSender interface {
	PushNotification(notificationText string) error
}

func NewPushNotificationSender() PushNotificationSender {
	return &pushNotificationSender{
		pushOverClient: pushover.New(config.PUSH_NOTIFICATION_ENV.PUSH_OVER_APP_TOKEN),
	}
}

func (pns *pushNotificationSender) PushNotification(notificationText string) error {
	recipient := pushover.NewRecipient(config.PUSH_NOTIFICATION_ENV.PUSH_OVER_APP_RECIPIENT)
	message := pushover.NewMessage(notificationText)

	response, err := pns.pushOverClient.SendMessage(message, recipient)
	if err != nil {
		return err
	}

	fmt.Println(response)
	return nil
}
