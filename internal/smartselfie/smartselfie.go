package smartselfie

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/nutcas3/smileid-go/internal/shared"
)

type AuthRequest struct {
	UserID      string `json:"user_id"`
	ImageBase64 string `json:"image_base64"`
	JobType     string `json:"job_type"`
}

type AuthResponse struct {
	Success      bool   `json:"success"`
	Verified     bool   `json:"verified"`
	UserID       string `json:"user_id"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type Service struct {
	Client *shared.BaseClient
}

func NewService(client *shared.BaseClient) *Service {
	return &Service{Client: client}
}

func (s *Service) AuthenticateSelfie(ctx context.Context, req AuthRequest) (*AuthResponse, error) {
	apiURL := s.Client.ResolveEndpoint("/smartselfie/authenticate")
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SmartSelfie request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	s.Client.SetSmileIDHeaders(httpReq)

	resp, err := s.Client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("SmartSelfie request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SmartSelfie API error: %s", resp.Status)
	}

	var ssResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&ssResp); err != nil {
		return nil, fmt.Errorf("failed to decode SmartSelfie response: %w", err)
	}
	return &ssResp, nil
}
