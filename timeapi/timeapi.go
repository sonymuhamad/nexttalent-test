package timeapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sonymuhamad/nexttalent-test/config"
	"github.com/sonymuhamad/nexttalent-test/dto"
	"io"
	"log"
	"net/http"
)

type TimeApi struct {
	client *http.Client
	cfg    config.EnvConfig
}

func NewTimeApiClient(cfg config.EnvConfig) *TimeApi {
	httpClient := &http.Client{}

	return &TimeApi{
		client: httpClient,
		cfg:    cfg,
	}
}

func (t *TimeApi) GetCurrentTimeByTimezone(_ context.Context, timeZone string) (dto.TimeApiResponse, error) {
	url := fmt.Sprintf("%s/api/Time/current/zone?timeZone=%s", t.cfg.TimeApiURL, timeZone)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return dto.TimeApiResponse{}, err
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return dto.TimeApiResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Request failed: %s\n", string(body))

		return dto.TimeApiResponse{}, errors.New("request failed")
	}

	var result dto.TimeApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return dto.TimeApiResponse{}, err
	}

	return result, nil
}
