package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"lynx/pkg/config"
)

var (
	ErrPermissionDenied = status.Error(
		codes.PermissionDenied,
		"permission denied",
	)
)

type Client interface {
	GetUserID(token string) (userID uuid.UUID, err error)
}

type DjangoAuthClient struct {
	cfg    config.AuthClientConfig
	client http.Client
}

func NewDjangoAuthClient(cfg config.AuthClientConfig) Client {
	return &DjangoAuthClient{
		cfg: cfg,
	}
}

func (d *DjangoAuthClient) GetUserID(token string) (userID uuid.UUID, err error) {
	url := url2.URL{
		Scheme: "http",
		Host:   d.cfg.Target,
		Path:   "/api/accounts/profile/my",
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

	resp, err := d.client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrPermissionDenied
		return
	}

	defer resp.Body.Close()



	respBody := struct {
		UserID string `json:"uuid"`
	}{}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&respBody)
	if err != nil {
		return
	}
	userID, err = uuid.FromString(respBody.UserID)
	return
}
