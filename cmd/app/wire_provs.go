package main

import (
	"net/http"

	"github.com/google/wire"

	"github.com/angelokurtis/go-talk/pkg/demo"
)

//nolint:unused // This function is used during compile-time to generate code for dependency injection
var providers = wire.NewSet(
	wire.Struct(new(http.Client)),
	wire.Struct(new(Manager), "*"),
	wire.Bind(new(demo.Showcase), new(*demo.MyShowcase)),
	demo.NewMyShowcase,
	demo.NewConfig,
	demo.NewElevenLabs,
)

type Manager struct {
	ElevenLabsAPI *demo.ElevenLabsAPI
}
