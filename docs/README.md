# SmileID Go SDK Documentation

Welcome to the SmileID Go SDK documentation. This section provides API reference and integration guides for all supported services.

## Table of Contents
- [Overview](#overview)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Service Reference](#service-reference)
  - [KYC](#kyc)
  - [Identity Verification](#identity-verification)
  - [Authentication](#authentication)
  - [Document Verification](#document-verification)
  - [SmartSelfie](#smartselfie)
  - [Job Status](#job-status)
  - [Business Verification](#business-verification)
  - [Supported ID Types](#supported-id-types)
- [Error Handling](#error-handling)
- [Examples](#examples)

---

## Overview
SmileID Go SDK enables seamless integration with Smile ID's digital KYC, identity verification, and authentication APIs.

## Getting Started
See the [README](../README.md) for installation and quickstart usage.

## Configuration
The SDK is configured via the `smileid.Config` struct:
```go
client := smileid.NewClient(smileid.Config{
    APIKey:    "your-api-key",
    PartnerID: "your-partner-id",
    Env:       "sandbox", // or "production"
})
```

## Service Reference

### KYC
```go
resp, err := client.KYC.VerifyUser(ctx, smileid.KYCRequest{ /* ... */ })
```

### Identity Verification
```go
resp, err := client.Identity.VerifyID(ctx, smileid.IdentityVerificationRequest{ /* ... */ })
```

### Authentication
```go
resp, err := client.Authentication.Authenticate(ctx, smileid.AuthRequest{ /* ... */ })
```

### Document Verification
```go
resp, err := client.DocumentVerification.VerifyDocument(ctx, smileid.DocumentVerificationRequest{ /* ... */ })
```

### SmartSelfie
```go
resp, err := client.SmartSelfie.Verify(ctx, smileid.SmartSelfieRequest{ /* ... */ })
```

### Job Status
```go
resp, err := client.JobStatus.GetStatus(ctx, smileid.JobStatusRequest{ /* ... */ })
```

### Business Verification
```go
resp, err := client.BusinessVerification.Verify(ctx, smileid.BusinessVerificationRequest{ /* ... */ })
```

### Supported ID Types
```go
resp, err := client.IDTypes.ListSupportedIDTypes(ctx, smileid.SupportedIDTypesRequest{ /* ... */ })
```

## Error Handling
All methods return detailed errors. Use Go's error wrapping/unwrapping to inspect and handle errors.

## Examples
See [../README.md](../README.md) for more code samples.
