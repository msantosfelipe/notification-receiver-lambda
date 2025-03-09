package lambda

import (
	"context"
	"encoding/json"
	"fmt"
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
	eventJSON, err := json.Marshal(event)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return err
	}

	var notification domain.Notification
	err = json.Unmarshal(eventJSON, &notification)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return err
	}

	if err := handler.notificationUc.ProcessNotification(notification); err != nil {
		fmt.Println("Error processing notification:", err)
		return err
	}

	return nil
}
