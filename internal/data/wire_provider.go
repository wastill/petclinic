//go:build wireinject
// +build wireinject

package data

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewEnt,
	NewData,
	NewRDBClient,
	NewUserRepo,
)
