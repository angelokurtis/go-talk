// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/angelokurtis/go-talk/pkg/demo"
)

// Injectors from wire_inj.go:

func NewManager() (*Manager, func(), error) {
	config, err := demo.NewConfig()
	if err != nil {
		return nil, nil, err
	}

	myShowcase := demo.NewMyShowcase(config)
	manager := &Manager{
		Example: myShowcase,
	}

	return manager, func() {
	}, nil
}
