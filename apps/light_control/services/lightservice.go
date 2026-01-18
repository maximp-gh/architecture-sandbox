package services

import (
	"encoding/json"
	"fmt"
	"lightcontrol/models"
	"net/http"
	"time"
)

// Service handles fetching light data from external API
type LightService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// LightResponse represents the response from the light API
type LightResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

// Creates a new light service
func NewLightService(baseURL string) *LightService {
	return &LightService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *LightService) CrudOps(sensorID string) (*models.Sensor, error) {
	url := "http://smarthome-app:8080/api/v1/sensors/" + sensorID
	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error finding sensor %s: %w", sensorID, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var devResp models.Sensor
	if err := json.NewDecoder(resp.Body).Decode(&devResp); err != nil {
		return nil, fmt.Errorf("error decoding light response: %w", err)
	}
	if devResp.Type != models.Light {
		return nil, fmt.Errorf("Wrong sensor type: %s", devResp.Type)
	}

	return &devResp, nil

}

// fetches light data for a specific location
func (s *LightService) GetLightByLocation(location string) (*LightResponse, error) {
	url := fmt.Sprintf("%s/light?location=%s", s.BaseURL, location)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching light data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var devResp LightResponse
	if err := json.NewDecoder(resp.Body).Decode(&devResp); err != nil {
		return nil, fmt.Errorf("error decoding light response: %w", err)
	}

	return &devResp, nil
}

// GetLightByID fetches light on/off status for a specific sensor ID
func (s *LightService) GetLightByID(sensorID string) (*LightResponse, error) {
	url := fmt.Sprintf("%s/light/%s", s.BaseURL, sensorID)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching light state: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var devResp LightResponse
	if err := json.NewDecoder(resp.Body).Decode(&devResp); err != nil {
		return nil, fmt.Errorf("error decoding light response: %w", err)
	}

	return &devResp, nil
}
