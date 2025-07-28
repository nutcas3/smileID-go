package identity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/nutcas3/smileid-go/internal/shared"
)

type VerificationRequest struct {
	CountryCode string `json:"country_code"`
	IDType      string `json:"id_type"`
	IDNumber    string `json:"id_number"`
}

type VerificationResponse struct {
	Success      bool   `json:"success"`
	Verified     bool   `json:"verified"`
	IDType       string `json:"id_type"`
	IDNumber     string `json:"id_number"`
	CountryCode  string `json:"country_code"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type Service struct {
	Client *shared.BaseClient
}

func NewService(client *shared.BaseClient) *Service {
	return &Service{Client: client}
}

func (s *Service) VerifyID(ctx context.Context, req VerificationRequest) (*VerificationResponse, error) {
	apiURL := s.Client.ResolveEndpoint("/identity/verify")
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal identity verification request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	s.Client.SetSmileIDHeaders(httpReq)

	resp, err := s.Client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("identity verification request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("identity verification API error: %s", resp.Status)
	}

	var idResp VerificationResponse
	if err := json.NewDecoder(resp.Body).Decode(&idResp); err != nil {
		return nil, fmt.Errorf("failed to decode identity verification response: %w", err)
	}
	return &idResp, nil
}
