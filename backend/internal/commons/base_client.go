package commons

import (
	"net/http"
	"time"
)

type BaseClient struct {
	BaseUrl    string
	HttpClient *http.Client
	Timeout    time.Duration
}
