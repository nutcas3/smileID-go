package business

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/nutcas3/smileid-go/internal/shared"
)

type VerificationRequest struct {
	CountryCode  string `json:"country_code"`
	BusinessID   string `json:"business_id"`
	BusinessType string `json:"business_type"`
}

type VerificationResponse struct {
	Success      bool   `json:"success"`
	Verified     bool   `json:"verified"`
	BusinessID   string `json:"business_id"`
	BusinessType string `json:"business_type"`
	CountryCode  string `json:"country_code"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type Service struct {
	Client *shared.BaseClient
}

func NewService(client *shared.BaseClient) *Service {
	return &Service{Client: client}
}

func (s *Service) VerifyBusiness(ctx context.Context, req VerificationRequest) (*VerificationResponse, error) {
	apiURL := s.Client.ResolveEndpoint("/business/verify")
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal business verification request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	s.Client.SetSmileIDHeaders(httpReq)

	resp, err := s.Client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("business verification request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("business verification API error: %s", resp.Status)
	}

	var bvResp VerificationResponse
	if err := json.NewDecoder(resp.Body).Decode(&bvResp); err != nil {
		return nil, fmt.Errorf("failed to decode business verification response: %w", err)
	}
	return &bvResp, nil
}
