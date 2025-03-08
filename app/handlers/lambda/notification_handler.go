package lambda

import (
	"context"
	"msantosfelipe/notification-receiver-lambda/domain"
)

type notificationHandler struct {
	notificationUc domain.NotificationUsecase
}

type NotificationHandler interface {
	ProcessNotification(ctx context.Context, event map[string]interface{}) error
}

func NewNotificationHandler(notificationUc domain.NotificationUsecase) NotificationHandler {
	return &notificationHandler{
		notificationUc: notificationUc,
	}
}

func (handler *notificationHandler) ProcessNotification(ctx context.Context, event map[string]interface{}) error {

}
