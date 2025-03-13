package usecase

import (
	"fmt"
	"msantosfelipe/notification-receiver-lambda/config"
	"msantosfelipe/notification-receiver-lambda/domain"
	"msantosfelipe/notification-receiver-lambda/infra"
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
		fmt.Sprintf("%s: - %s", notification.AppName, notification.Body),
	)
}

func validateNotification(notification *domain.Notification) error {
	for _, i := range config.ENV.APPS_ALLOWED {
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
