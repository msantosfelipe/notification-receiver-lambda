package usecase

import (
	"fmt"
	"msantosfelipe/notification-receiver-lambda/internal/config"
	"msantosfelipe/notification-receiver-lambda/internal/domain"
	"msantosfelipe/notification-receiver-lambda/internal/infra"
	"slices"
)

type usecase struct {
	pushNotificationSender infra.PushNotificationSender
}

func NewNotificationUsecase(
	pushNotificationSender infra.PushNotificationSender,
) domain.NotificationUsecase {
	return &usecase{
		pushNotificationSender: pushNotificationSender,
	}
}

func (uc *usecase) ProcessNotification(notification *domain.Notification) error {
	if err := validateNotification(notification); err != nil {
		fmt.Println("Error validating notification:", err)
		return err
	}

	return uc.pushNotificationSender.PushNotification(
		fmt.Sprintf("%s: %s - %s", notification.AppName, notification.Title, notification.Body),
	)
}

func validateNotification(notification *domain.Notification) error {
	for _, i := range config.ENV.AppsAllowed {
		if i.App == notification.AppName {
			if i.FullyAllowed {
				return nil
			}
			if slices.Contains(i.AllowedTitles, notification.Title) {
				return nil
			}
		}
	}
	return fmt.Errorf("invalid app name: %s", notification.AppName)
}
