package smileid

import (
	"net/http"
	"time"

	"github.com/nutcas3/smileid-go/internal/kyc"
	"github.com/nutcas3/smileid-go/internal/identity"
	"github.com/nutcas3/smileid-go/internal/auth"
	"github.com/nutcas3/smileid-go/internal/document"
	"github.com/nutcas3/smileid-go/internal/smartselfie"
	"github.com/nutcas3/smileid-go/internal/jobstatus"
	"github.com/nutcas3/smileid-go/internal/business"
	"github.com/nutcas3/smileid-go/internal/idtypes"
	"github.com/nutcas3/smileid-go/internal/shared"
)


 type Config struct {
	APIKey    string
	PartnerID string
	Env       string // "sandbox" or "production"
	Timeout   time.Duration // Optional: request timeout
}

// Client is the main struct for interacting with SmileID services
 type Client struct {
	config Config
	httpClient *http.Client
	baseClient *shared.BaseClient

	KYC           *kyc.Service
	Identity      *identity.Service
	Authentication *auth.Service
	DocumentVerification *document.Service
	SmartSelfie *smartselfie.Service
	JobStatus *jobstatus.Service
	BusinessVerification *business.Service
	IDTypes *idtypes.Service
}

// NewClient creates a new SmileID client
func NewClient(cfg Config) *Client {
	if cfg.Timeout == 0 {
		cfg.Timeout = 15 * time.Second
	}
	hc := &http.Client{Timeout: cfg.Timeout}
	base := &shared.BaseClient{
		Config: shared.Config{
			APIKey:    cfg.APIKey,
			PartnerID: cfg.PartnerID,
			Env:       cfg.Env,
		},
		HTTPClient: hc,
	}
	c := &Client{
		config: cfg,
		httpClient: hc,
		baseClient: base,
	}
	c.KYC = kyc.NewService(base)
	c.Identity = identity.NewService(base)
	c.Authentication = auth.NewService(base)
	c.DocumentVerification = document.NewService(base)
	c.SmartSelfie = smartselfie.NewService(base)
	c.JobStatus = jobstatus.NewService(base)
	c.BusinessVerification = business.NewService(base)
	c.IDTypes = idtypes.NewService(base)
	return c
}
