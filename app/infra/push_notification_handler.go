package infra

import (
	"log"
	"msantosfelipe/notification-receiver-lambda/config"

	"github.com/gregdel/pushover"
)

type pushNotificationSender struct {
	pushOverClient *pushover.Pushover
}

type PushNotificationSender interface {
	PushNotification(notificationText string)
}

func NewPushNotificationSender() PushNotificationSender {
	return &pushNotificationSender{
		pushOverClient: pushover.New(config.PUSH_NOTIFICATION_ENV.PUSH_OVER_APP_TOKEN),
	}
}

func (pns *pushNotificationSender) PushNotification(notificationText string) {
	pns.PushOverlPushNotification(notificationText)
}

func (pns *pushNotificationSender) PushOverlPushNotification(notificationText string) {
	recipient := pushover.NewRecipient(config.PUSH_NOTIFICATION_ENV.PUSH_OVER_APP_RECIPIENT)
	message := pushover.NewMessage(notificationText)

	response, err := pns.pushOverClient.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	log.Println(response)
}
