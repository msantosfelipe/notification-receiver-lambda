package infra

import (
	"fmt"
	"msantosfelipe/notification-receiver-lambda/internal/config"

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
		pushOverClient: pushover.New(config.PushNotificationEnv.PushOverAppToken),
	}
}

func (pns *pushNotificationSender) PushNotification(notificationText string) error {
	recipient := pushover.NewRecipient(config.PushNotificationEnv.PushOverAppRecipient)
	message := pushover.NewMessage(notificationText)

	response, err := pns.pushOverClient.SendMessage(message, recipient)
	if err != nil {
		return err
	}

	fmt.Println(response)
	return nil
}
