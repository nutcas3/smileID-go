package idtypes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/nutcas3/smileid-go/internal/shared"

	
)

type ListRequest struct {
	CountryCode string `json:"country_code"`
}

type ListResponse struct {
	Success      bool     `json:"success"`
	IDTypes      []string `json:"id_types"`
	ErrorMessage string   `json:"error_message,omitempty"`
}

type Service struct {
	Client *shared.BaseClient
}

func NewService(client *shared.BaseClient) *Service {
	return &Service{Client: client}
}

func (s *Service) ListSupportedIDTypes(ctx context.Context, req ListRequest) (*ListResponse, error) {
	apiURL := s.Client.ResolveEndpoint("/id-types/list")
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal supported ID types request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	s.Client.SetSmileIDHeaders(httpReq)

	resp, err := s.Client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("supported ID types request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("supported ID types API error: %s", resp.Status)
	}

	var idResp ListResponse
	if err := json.NewDecoder(resp.Body).Decode(&idResp); err != nil {
		return nil, fmt.Errorf("failed to decode supported ID types response: %w", err)
	}
	return &idResp, nil
}
