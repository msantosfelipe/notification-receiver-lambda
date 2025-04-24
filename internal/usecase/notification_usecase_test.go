package usecase

import (
	"fmt"
	"msantosfelipe/notification-receiver-lambda/config"
	"msantosfelipe/notification-receiver-lambda/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPushNotificationSender struct {
	mock.Mock
}

func (m *MockPushNotificationSender) PushNotification(notificationText string) error {
	args := m.Called(notificationText)
	return args.Error(0)
}

func TestProcessNotification_Success(t *testing.T) {
	mockSender := new(MockPushNotificationSender)
	uc := NewNotificationUsecase(mockSender)

	notification := domain.Notification{
		AppName: "XP",
		Title:   "Unit test",
		Body:    "new notification arrived",
	}

	config.ENV.APPS_ALLOWED = []domain.AppAllowed{
		{
			App:          "XP",
			FullyAllowed: true,
			AllowedTitles: []string{
				"Unit test",
			},
		},
	}

	mockSender.On("PushNotification", "XP: Unit test - new notification arrived").Return(nil)

	err := uc.ProcessNotification(&notification)

	assert.NoError(t, err)
	mockSender.AssertExpectations(t)
}

func TestProcessNotification_InvalidApp(t *testing.T) {
	mockSender := new(MockPushNotificationSender)
	uc := NewNotificationUsecase(mockSender)

	notification := domain.Notification{
		AppName: "Invalid",
		Title:   "Unit test",
		Body:    "new notification arrived",
	}

	config.ENV.APPS_ALLOWED = []domain.AppAllowed{
		{
			App:          "XP",
			FullyAllowed: true,
			AllowedTitles: []string{
				"Unit test",
			},
		},
	}

	err := uc.ProcessNotification(&notification)

	assert.Error(t, err)
	assert.EqualError(t, fmt.Errorf("invalid app name: Invalid"), err.Error())
	mockSender.AssertExpectations(t)
}
