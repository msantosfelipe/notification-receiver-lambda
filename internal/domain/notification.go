package domain

type Notification struct {
	AppName string `json:"app"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type NotificationUsecase interface {
	ProcessNotification(notification *Notification) error
}
