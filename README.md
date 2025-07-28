# SmileID Go SDK

SmileID Go SDK provides easy integration for Smile ID's digital identity verification, KYC, fraud prevention, and user authentication services for Africa.

## Features
- Digital KYC: Instant user onboarding and verification
- Identity Verification: Support for 8500+ ID types, 220+ countries
- User Authentication: Secure, real-time authentication
- Seamless Integration: Simple, idiomatic Go API

## Installation
```sh
go get github.com/nutcas3/smileid-go
```

## Usage Example

First, import the package from the root:
```go
import (
    "fmt"
    "github.com/nutcas3/smileid-go"
)
```

### Digital KYC
```go
ctx := context.Background()
client := smileid.NewClient(smileid.Config{
    APIKey:    "your-api-key",
    PartnerID: "your-partner-id",
    Env:       smileid.Sandbox,
})
kycReq := smileid.KYCRequest{
    CountryCode: "NG",
    IDType:      "NIN",
    IDNumber:    "12345678901",
    FirstName:   "Jane",
    LastName:    "Doe",
}
kycResp, err := client.KYC.VerifyUser(ctx, kycReq)
if err != nil {
    // Handle error
    log.Fatal(err)
}
fmt.Println(kycResp)
```

### Identity Verification
```go
idReq := smileid.IdentityVerificationRequest{
    CountryCode: "GH",
    IDType:      "PASSPORT",
    IDNumber:    "A1234567",
}
idResp, err := client.Identity.VerifyID(ctx, idReq)
if err != nil {
    log.Fatal(err)
}
fmt.Println(idResp)
```

### User Authentication
```go
authReq := smileid.AuthRequest{
    UserID:   "user-123",
    Event:    "login",
    DeviceID: "device-abc",
}
authResp, err := client.Authentication.Authenticate(ctx, authReq)
if err != nil {
    log.Fatal(err)
}
fmt.Println(authResp)
```

## Error Handling
All API methods return detailed errors. Use Go's error wrapping/unwrapping to inspect and handle errors.

---

## Documentation
See [docs](./docs/) for full API reference and integration guides.

---

© 2025 Smile ID. All rights reserved.
