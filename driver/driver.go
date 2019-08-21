package driver

import (
	"github.com/tyaps/hydra/driver/configuration"
)

type Driver interface {
	Configuration() configuration.Provider
	Registry() Registry
	CallRegistry() Driver
}
