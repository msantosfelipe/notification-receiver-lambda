package usecase

import (
	"fmt"
	"msantosfelipe/notification-receiver-lambda/config"
	"msantosfelipe/notification-receiver-lambda/domain"
	"msantosfelipe/notification-receiver-lambda/infra"
	"strings"
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

func (uc *usecase) ProcessNotification(notification domain.Notification) {
	if err := validateNotification(notification); err != nil {
		fmt.Println(err)
		return
	}

	uc.pushNotificationSender.PushNotification(
		fmt.Sprintf("%s: - %s", notification.Name, notification.Body),
	)
}

func validateNotification(notification domain.Notification) error {
	if err := isValidApp(notification.Name); err != nil {
		return err
	}
	if err := isValidTitle(notification.Name, notification.Title); err != nil {
		return err
	}
	return nil
}

func isValidApp(appName string) error {
	if config.ENV.ALLOW_ALL_APPS {
		return nil
	}

	for _, i := range config.ENV.ALLOWED_APPS {
		if i == appName {
			return nil
		}
	}
	return fmt.Errorf("invalid app name: %s", appName)
}

func isValidTitle(appName, title string) error {
	if config.ENV.ALLOW_ALL_APPS {
		return nil
	}

	for _, i := range config.ENV.ALLOWED_TITLES {
		if strings.Contains(i, title) {
			return nil
		}
	}
	return fmt.Errorf("invalid title name: %s of app %s", title, appName)
}
