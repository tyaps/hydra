package jwk

import (
	"github.com/tyaps/hydra/driver/configuration"
	"github.com/tyaps/hydra/x"
)

type InternalRegistry interface {
	x.RegistryWriter
	x.RegistryLogger
	Registry
}

type Registry interface {
	KeyManager() Manager
	KeyGenerators() map[string]KeyGenerator
	KeyCipher() *AEAD
}

type Configuration interface {
	configuration.Provider
}
