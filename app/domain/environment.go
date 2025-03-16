package domain

type Config struct {
	IS_LOCAL       bool
	VALID_API_KEY  string
	ALLOW_ALL_APPS bool
	APPS_ALLOWED   []AppAllowed
	ALLOWED_APPS   []string
	ALLOWED_TITLES []string
}

type AppAllowed struct {
	App           string   `json:"app"`
	FullyAllowed  bool     `json:"fully_allowed"`
	AllowedTitles []string `json:"allowed_titles"`
}

type PushNotification struct {
	PUSH_OVER_APP_TOKEN     string
	PUSH_OVER_APP_RECIPIENT string
}
