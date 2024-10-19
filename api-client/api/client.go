package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL = "https://botw-compendium.herokuapp.com/api/v3/compendium"

type HttpCilent interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	baseUrl    string
	httpCilent HttpCilent
}

func NewClient(baseUrl string, httpClient HttpCilent) *Client {
	return &Client{
		baseUrl:    baseUrl,
		httpCilent: httpClient,
	}
}

type GetMonsterResponse struct {
	Data    []Monster `json:"data"`
	Message string    `json:"message"`
	Status  int64     `json:"status"`
}

type Monster struct {
	Category        string   `json:"category"`
	CommonLocations []string `json:"common_locations"`
	Description     string   `json:"description"`
	Dlc             bool     `json:"dlc"`
	Drops           []string `json:"drops"`
	ID              int64    `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
}

func (c *Client) GetMonster() (*GetMonsterResponse, error) {
	req, err := http.NewRequest("GET", c.baseUrl+"/category/monsters", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create monster request:%w", err)
	}

	reqUrl := req.URL
	queryParams := req.URL.Query()
	queryParams.Set("game", "totk")
	reqUrl.RawQuery = queryParams.Encode()

	resp, err := c.httpCilent.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to submit monster http request:%w", err)
	}

	var response *GetMonsterResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal monster http response:%w", err)
	}

	return response, nil
}
