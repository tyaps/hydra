package oauth2

import (
	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/openid"
	"github.com/tyaps/hydra/client"
	"github.com/tyaps/hydra/consent"
	"github.com/tyaps/hydra/driver/configuration"
	"github.com/tyaps/hydra/jwk"
	"github.com/tyaps/hydra/x"
)

type InternalRegistry interface {
	client.Registry
	x.RegistryWriter
	x.RegistryLogger
	consent.Registry
	Registry
}

type Registry interface {
	OAuth2Storage() x.FositeStorer
	OAuth2Provider() fosite.OAuth2Provider
	AudienceStrategy() fosite.AudienceMatchingStrategy
	ScopeStrategy() fosite.ScopeStrategy

	AccessTokenJWTStrategy() jwk.JWTStrategy
	OpenIDJWTStrategy() jwk.JWTStrategy

	OpenIDConnectRequestValidator() *openid.OpenIDConnectRequestValidator
}

type Configuration interface {
	configuration.Provider
}
