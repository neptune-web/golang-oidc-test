package storage

import (
	"time"

	"github.com/zitadel/oidc/pkg/oidc"
	"github.com/zitadel/oidc/pkg/op"
)

var (
	// we use the default login UI and pass the (auth request) id
	defaultLoginURL = func(id string) string {
		return "/login/?authRequestID=" + id
	}

	// clients to be used by the storage interface
	clients = map[string]*Client{}
)

// Client represents the storage model of an OAuth/OIDC client
// this could also be your database model
type Client struct {
	id                             string
	secret                         string
	redirectURIs                   []string
	postLogoutRedirectURIs         []string
	applicationType                op.ApplicationType
	authMethod                     oidc.AuthMethod
	loginURL                       func(string) string
	responseTypes                  []oidc.ResponseType
	grantTypes                     []oidc.GrantType
	accessTokenType                op.AccessTokenType
	devMode                        bool
	idTokenUserinfoClaimsAssertion bool
	clockSkew                      time.Duration
}

// GetID must return the client_id
func (c *Client) GetID() string {
	return c.id
}

// RedirectURIs must return the registered redirect_uris for Code and Implicit Flow
func (c *Client) RedirectURIs() []string {
	return c.redirectURIs
}

// PostLogoutRedirectURIs must return the registered post_logout_redirect_uris for sign-outs
func (c *Client) PostLogoutRedirectURIs() []string {
	return c.postLogoutRedirectURIs
}

// ApplicationType must return the type of the client (app, native, user agent)
func (c *Client) ApplicationType() op.ApplicationType {
	return c.applicationType
}

// AuthMethod must return the authentication method (client_secret_basic, client_secret_post, none, private_key_jwt)
func (c *Client) AuthMethod() oidc.AuthMethod {
	return c.authMethod
}

// ResponseTypes must return all allowed response types (code, id_token token, id_token)
// these must match with the allowed grant types
func (c *Client) ResponseTypes() []oidc.ResponseType {
	return c.responseTypes
}

// GrantTypes must return all allowed grant types (authorization_code, refresh_token, urn:ietf:params:oauth:grant-type:jwt-bearer)
func (c *Client) GrantTypes() []oidc.GrantType {
	return c.grantTypes
}

// LoginURL will be called to redirect the user (agent) to the login UI
// you could implement some logic here to redirect the users to different login UIs depending on the client
func (c *Client) LoginURL(id string) string {
	return c.loginURL(id)
}

// AccessTokenType must return the type of access token the client uses (Bearer (opaque) or JWT)
func (c *Client) AccessTokenType() op.AccessTokenType {
	return c.accessTokenType
}

// IDTokenLifetime must return the lifetime of the client's id_tokens
func (c *Client) IDTokenLifetime() time.Duration {
	return 1 * time.Hour
}

// DevMode enables the use of non-compliant configs such as redirect_uris (e.g. http schema for user agent client)
func (c *Client) DevMode() bool {
	return c.devMode
}

// RestrictAdditionalIdTokenScopes allows specifying which custom scopes shall be asserted into the id_token
func (c *Client) RestrictAdditionalIdTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

// RestrictAdditionalAccessTokenScopes allows specifying which custom scopes shall be asserted into the JWT access_token
func (c *Client) RestrictAdditionalAccessTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

// IsScopeAllowed enables Client specific custom scopes validation
// in this example we allow the CustomScope for all clients
func (c *Client) IsScopeAllowed(scope string) bool {
	return scope == CustomScope
}

// IDTokenUserinfoClaimsAssertion allows specifying if claims of scope profile, email, phone and address are asserted into the id_token
// even if an access token if issued which violates the OIDC Core spec
//(5.4. Requesting Claims using Scope Values: https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims)
// some clients though require that e.g. email is always in the id_token when requested even if an access_token is issued
func (c *Client) IDTokenUserinfoClaimsAssertion() bool {
	return c.idTokenUserinfoClaimsAssertion
}

// ClockSkew enables clients to instruct the OP to apply a clock skew on the various times and expirations
//(subtract from issued_at, add to expiration, ...)
func (c *Client) ClockSkew() time.Duration {
	return c.clockSkew
}

// RegisterClients enables you to register clients for the example implementation
// there are some clients (web and native) to try out different cases
// add more if necessary
//
// RegisterClients should be called before the Storage is used so that there are
// no race conditions.
func RegisterClients(registerClients ...*Client) {
	for _, client := range registerClients {
		clients[client.id] = client
	}
}

// TODO: move clients to DB
// js and api can use the same client, no need to split them, the ones below are just for testing purposes

// Test client for js app
func WebClient(id, secret string) *Client {
	return &Client{
		id:     id,
		secret: secret,
		redirectURIs: []string{
			"http://localhost:15000/oidc-client-sample.html",
			"http://localhost:15000/code-identityserver-sample.html",
			"http://localhost:15000/code-identityserver-sample-silent.html",
			"http://localhost:15000/code-identityserver-sample-popup-signin.html",
		},
		postLogoutRedirectURIs: []string{
			"http://localhost:15000/oidc-client-sample.html",
			"http://localhost:15000/code-identityserver-sample.html",
			"http://localhost:15000/code-identityserver-sample-popup-signout.html",
		},
		applicationType:                op.ApplicationTypeWeb,
		authMethod:                     oidc.AuthMethodNone,
		loginURL:                       defaultLoginURL,
		responseTypes:                  []oidc.ResponseType{oidc.ResponseTypeIDToken, oidc.ResponseTypeCode},
		grantTypes:                     oidc.AllGrantTypes, //[]oidc.GrantType{oidc.GrantTypeCode, oidc.GrantTypeRefreshToken},
		accessTokenType:                0,
		devMode:                        true,
		idTokenUserinfoClaimsAssertion: false,
		clockSkew:                      0,
	}
}

// Test client for go app
func WebClient2(id, secret string) *Client {
	return &Client{
		id:                             id,
		secret:                         secret,
		redirectURIs:                   []string{"http://localhost:9999/auth/callback"},
		applicationType:                op.ApplicationTypeWeb,
		authMethod:                     oidc.AuthMethodBasic,
		loginURL:                       defaultLoginURL,
		responseTypes:                  []oidc.ResponseType{oidc.ResponseTypeCode},
		grantTypes:                     []oidc.GrantType{oidc.GrantTypeCode, oidc.GrantTypeRefreshToken},
		accessTokenType:                0,
		devMode:                        true,
		idTokenUserinfoClaimsAssertion: false,
		clockSkew:                      0,
	}
}
