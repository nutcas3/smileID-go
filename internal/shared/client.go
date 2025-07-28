package shared

import (
	"net/http"
	"github.com/nutcas3/smileid-go/internal/signature"
)

type Config struct {
	APIKey    string
	PartnerID string
	Env       string
}

type BaseClient struct {
	Config     Config
	HTTPClient *http.Client
}

func (c *BaseClient) ResolveEndpoint(path string) string {
	base := "https://api.smileid.com/v1"
	return base + path
}

func (c *BaseClient) SetSmileIDHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Key", c.Config.APIKey)
	req.Header.Set("Partner-ID", c.Config.PartnerID)
	sig, ts, _ := signature.GenerateSignature(c.Config.APIKey, c.Config.PartnerID)
	req.Header.Set("X-Signature", sig)
	req.Header.Set("X-Timestamp", ts)
}
