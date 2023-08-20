package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"

	"github.com/remustaev/telegram-bot/internal/entity"
)

// authz login password

// OpenWeatherMap API
// https://api.openweathermap.org/data/2.5/weather?lat=52.3417853&lon=4.9040615&appid=f2be13b8fe65a7e49e1b4314d5404af8

type OpenWeatherMapClient struct {
	apiToken string
	client   *http.Client

	cache *lru.Cache[entity.Location, entity.Weather]
}

func New(apiKey string) *OpenWeatherMapClient {
	// todo
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	cache, _ := lru.New[entity.Location, entity.Weather](128)

	return &OpenWeatherMapClient{
		apiToken: apiKey,
		client:   client,
		cache:    cache,
	}
}

func (c *OpenWeatherMapClient) GetWeather(ctx context.Context, location entity.Location, _ time.Time) (entity.Weather, error) {
	w, ok := c.cache.Get(location)
	if ok {
		return w, nil
	}

	// api doc - https://openweathermap.org/current
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.openweathermap.org/data/2.5/weather", nil)
	if err != nil {
		return entity.Weather{}, fmt.Errorf("create OpenWeatherMapClient.GetWeather request failed %w: %w", entity.ErrInternal, err)
	}

	q := req.URL.Query()
	q.Add("lat", fmt.Sprintf("%f", location.Latitude))
	q.Add("lon", fmt.Sprintf("%f", location.Longitude))
	q.Add("units", "metric")
	q.Add("appid", c.apiToken)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return entity.Weather{}, fmt.Errorf("execute OpenWeatherMapClient.GetWeather request failed %w: %w", entity.ErrInternal, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entity.Weather{}, fmt.Errorf("execute OpenWeatherMapClient.GetWeather request failed %w: %s", entity.ErrInternal, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entity.Weather{}, fmt.Errorf("read OpenWeatherMapClient.GetWeather response body failed %w: %w", entity.ErrInternal, err)
	}

	var respDTO ResponseDTO
	err = json.Unmarshal(body, &respDTO)
	if err != nil {
		return entity.Weather{}, fmt.Errorf("read OpenWeatherMapClient.GetWeather response body failed %w: %w", entity.ErrInternal, err)
	}

	w = mapToWeather(respDTO)
	c.cache.Add(location, w)

	return w, nil
}
