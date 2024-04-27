//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func NewManager() (*Manager, func(), error) {
	wire.Build(providers)
	return nil, nil, nil
}
