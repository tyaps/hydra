package consent

import (
	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/openid"
	"github.com/tyaps/hydra/client"
	"github.com/tyaps/hydra/driver/configuration"
	"github.com/tyaps/hydra/jwk"
	"github.com/tyaps/hydra/x"
)

type InternalRegistry interface {
	x.RegistryWriter
	x.RegistryCookieStore
	x.RegistryLogger
	Registry
	client.Registry

	OAuth2Storage() x.FositeStorer
	OpenIDJWTStrategy() jwk.JWTStrategy
	OpenIDConnectRequestValidator() *openid.OpenIDConnectRequestValidator
	ScopeStrategy() fosite.ScopeStrategy
}

type Registry interface {
	ConsentManager() Manager
	ConsentStrategy() Strategy
	SubjectIdentifierAlgorithm() map[string]SubjectIdentifierAlgorithm
}

type Configuration interface {
	configuration.Provider
}
