package jobstatus

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/nutcas3/smileid-go/internal/shared"
)

type StatusRequest struct {
	JobID   string `json:"job_id"`
	UserID  string `json:"user_id"`
}

type StatusResponse struct {
	Success      bool   `json:"success"`
	Status       string `json:"status"`
	JobID        string `json:"job_id"`
	UserID       string `json:"user_id"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type Service struct {
	Client *shared.BaseClient
}

func NewService(client *shared.BaseClient) *Service {
	return &Service{Client: client}
}

func (s *Service) GetJobStatus(ctx context.Context, req StatusRequest) (*StatusResponse, error) {
	apiURL := s.Client.ResolveEndpoint("/job/status")
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal job status request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	s.Client.SetSmileIDHeaders(httpReq)

	resp, err := s.Client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("job status request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("job status API error: %s", resp.Status)
	}

	var jsResp StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&jsResp); err != nil {
		return nil, fmt.Errorf("failed to decode job status response: %w", err)
	}
	return &jsResp, nil
}
