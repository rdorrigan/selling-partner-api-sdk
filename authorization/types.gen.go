// Package authorization provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package authorization

// AuthorizationCode defines model for AuthorizationCode.
type AuthorizationCode struct {

	// A Login with Amazon (LWA) authorization code that can be exchanged for a refresh token and access token that authorize you to make calls to a Selling Partner API.
	AuthorizationCode *string `json:"authorizationCode,omitempty"`
}

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList []Error

// GetAuthorizationCodeResponse defines model for GetAuthorizationCodeResponse.
type GetAuthorizationCodeResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// A Login with Amazon (LWA) authorization code.
	Payload *AuthorizationCode `json:"payload,omitempty"`
}

// GetAuthorizationCodeParams defines parameters for GetAuthorizationCode.
type GetAuthorizationCodeParams struct {

	// The seller ID of the seller for whom you are requesting Selling Partner API authorization. This must be the seller ID of the seller who authorized your application on the Marketplace Appstore.
	SellingPartnerId string `json:"sellingPartnerId"`

	// Your developer ID. This must be one of the developer ID values that you provided when you registered your application in Developer Central.
	DeveloperId string `json:"developerId"`

	// The MWS Auth Token that was generated when the seller authorized your application on the Marketplace Appstore.
	MwsAuthToken string `json:"mwsAuthToken"`
}
