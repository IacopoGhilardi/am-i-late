package tomtom

import (
	"net/http"
	"time"

	"github.com/iacopoGhilardi/amILate/internal/commons"
)

const BaseUrl = "https://api.tomtom.com"

type Client struct {
	commons.BaseClient
	ApiKey string
}

func NewClient(apiKey string, timeout time.Duration) *Client {
	return &Client{
		BaseClient: commons.BaseClient{
			BaseUrl: BaseUrl,
			HttpClient: &http.Client{
				Timeout: timeout,
			},
		},
		ApiKey: apiKey,
	}
}
