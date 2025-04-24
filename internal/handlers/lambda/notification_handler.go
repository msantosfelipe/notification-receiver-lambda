package lambda

import (
	"encoding/json"
	"errors"
	"fmt"
	"msantosfelipe/notification-receiver-lambda/internal/config"
	"msantosfelipe/notification-receiver-lambda/internal/domain"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

const ApiKeyValidationHeader = "apikey"

type notificationHandler struct {
	notificationUc domain.NotificationUsecase
}

type NotificationHandler interface {
	ProcessNotification(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func NewNotificationHandler(notificationUc domain.NotificationUsecase) NotificationHandler {
	return &notificationHandler{
		notificationUc: notificationUc,
	}
}

func (handler *notificationHandler) ProcessNotification(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !isAuthorized(request.Headers) {
		fmt.Println("Unauthorized - invalid apiKey in headers: ", request.Headers)
		return errorResponse(http.StatusUnauthorized, errors.New("is unauthorized"))
	}

	notification, err := extractNotificationBody(request.Body)
	if err != nil {
		fmt.Println("Error extracting notification body:", err)
		return errorResponse(http.StatusBadRequest, err)
	}

	if err := handler.notificationUc.ProcessNotification(notification); err != nil {
		fmt.Println("Error processing notification:", err)
		return errorResponse(http.StatusInternalServerError, err)
	}

	return events.APIGatewayProxyResponse{
		Body:       "OK",
		StatusCode: http.StatusOK,
	}, nil
}

func errorResponse(statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       err.Error(),
	}, nil
}

func isAuthorized(headers map[string]string) bool {
	apiKeyHeader := headers[ApiKeyValidationHeader]
	if apiKeyHeader != "" {
		return apiKeyHeader == config.ENV.ValidApiKey
	}
	return false
}

func extractNotificationBody(body string) (*domain.Notification, error) {
	bodyBytes := []byte(body)
	var notification domain.Notification
	err := json.Unmarshal(bodyBytes, &notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}
