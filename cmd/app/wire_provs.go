package main

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/google/wire"

	"github.com/angelokurtis/go-talk/internal/openai"
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
	openai.NewConfig,
	openai.NewClient,
)

type Manager struct {
	OpenAPI *azopenai.Client
}
