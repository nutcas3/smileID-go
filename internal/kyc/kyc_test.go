package kyc

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/nutcas3/smileid-go/internal/shared"
)

func TestVerifyUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":true,"verified":true,"full_name":"Jane Doe","id_number":"12345678901","id_type":"NIN","country_code":"NG","dob":"1990-01-01"}`))
	}))
	defer ts.Close()

	client := &shared.BaseClient{
		Config: shared.Config{APIKey: "test", PartnerID: "pid", Env: "sandbox"},
		HTTPClient: ts.Client(),
	}
	service := NewService(client)
	// Patch the service to use the test endpoint
	service.EndpointOverride = ts.URL

	resp, err := service.VerifyUser(context.Background(), KYCRequest{
		CountryCode: "NG",
		IDType:      "NIN",
		IDNumber:    "12345678901",
		FirstName:   "Jane",
		LastName:    "Doe",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.Success || !resp.Verified || resp.FullName != "Jane Doe" {
		t.Errorf("unexpected response: %+v", resp)
	}
}
