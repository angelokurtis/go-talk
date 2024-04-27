package main

import (
	"github.com/google/wire"

	"github.com/angelokurtis/go-talk/pkg/demo"
)

//nolint:unused // This function is used during compile-time to generate code for dependency injection
var providers = wire.NewSet(
	wire.Struct(new(Manager), "*"),
	wire.Bind(new(demo.Showcase), new(*demo.MyShowcase)),
	demo.NewMyShowcase,
	demo.NewConfig,
)

type Manager struct {
	Example demo.Showcase
}
