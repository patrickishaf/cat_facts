package lib

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type httpClient struct {
	client http.Client
}

func NewHttpClient() *httpClient {
	return &httpClient{
		client: http.Client{
			Timeout: time.Second * time.Second * time.Duration(config.RequestTimeoutSeconds),
		},
	}
}

func (hc *httpClient) Get(url string) (*[]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := hc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if strings.Split(res.Status, "")[0] != "2" {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf(string(b))
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
