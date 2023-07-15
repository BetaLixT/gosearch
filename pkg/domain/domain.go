// Package domain containing business logic, contracts and models
package domain

import (
	"github.com/google/wire"
	"github.com/BetaLixT/gex/pkg/domain/usecases"
)

// DependencySet dependencies provided by the domain
var DependencySet = wire.NewSet(
	usecases.NewUseCases,
)
