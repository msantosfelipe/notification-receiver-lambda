package domain

type Config struct {
	IsLocal      bool
	ValidApiKey  string
	AllowAllApps bool
	AppsAllowed  []AppAllowed
}

type AppAllowed struct {
	App           string   `json:"app"`
	FullyAllowed  bool     `json:"fully_allowed"`
	AllowedTitles []string `json:"allowed_titles"`
}

type PushNotification struct {
	PushOverAppToken     string
	PushOverAppRecipient string
}
