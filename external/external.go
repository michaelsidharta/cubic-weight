package external

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/michaelsidharta/cubic-weight/entity"
)

var baseAPIURL string

type IAPI interface {
	Get(ctx context.Context, URL string) (entity.APIResponse, error)
}

type API struct {
	baseAPIURL string
}

func Init(apiURL string) IAPI {
	return &API{baseAPIURL: apiURL}
}

func (a API) Get(ctx context.Context, URL string) (entity.APIResponse, error) {
	if URL == "" {
		return entity.APIResponse{}, errors.New("Empty URL")
	}
	url := a.baseAPIURL + URL
	resp, err := http.Get(url)
	if err != nil {
		return entity.APIResponse{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.APIResponse{}, err
	}

	var response entity.APIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return entity.APIResponse{}, err
	}
	return response, nil
}
