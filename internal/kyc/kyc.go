package kyc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/nutcas3/smileid-go/internal/shared"
)

type KYCRequest struct {
	CountryCode string `json:"country_code"`
	IDType      string `json:"id_type"`
	IDNumber    string `json:"id_number"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	DOB         string `json:"dob,omitempty"`
}

type KYCResponse struct {
	Success      bool   `json:"success"`
	Verified     bool   `json:"verified"`
	FullName     string `json:"full_name"`
	IDNumber     string `json:"id_number"`
	IDType       string `json:"id_type"`
	CountryCode  string `json:"country_code"`
	DOB          string `json:"dob"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type Service struct {
	Client           *shared.BaseClient
	EndpointOverride string // for testing
}

func NewService(client *shared.BaseClient) *Service {
	return &Service{Client: client}
}

func (s *Service) VerifyUser(ctx context.Context, req KYCRequest) (*KYCResponse, error) {
	apiURL := s.Client.ResolveEndpoint("/kyc/verify")
	if s.EndpointOverride != "" {
		apiURL = s.EndpointOverride
	}
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal KYC request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	s.Client.SetSmileIDHeaders(httpReq)

	resp, err := s.Client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("KYC request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("KYC API error: %s", resp.Status)
	}

	var kycResp KYCResponse
	if err := json.NewDecoder(resp.Body).Decode(&kycResp); err != nil {
		return nil, fmt.Errorf("failed to decode KYC response: %w", err)
	}
	return &kycResp, nil
}
