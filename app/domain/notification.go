package domain

type Notification struct {
	Name  string `json:"app"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NotificationUsecase interface {
	ProcessNotification(notification Notification)
}
