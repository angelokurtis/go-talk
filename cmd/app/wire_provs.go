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
	demo.NewConfig,
	demo.NewElevenLabs,
	demo.NewMP3Writer,
	demo.NewMyShowcase,
	demo.NewSpeechGenerator,
	openai.NewClient,
	openai.NewConfig,
	wire.Bind(new(demo.OpenAI), new(*azopenai.Client)),
	wire.Bind(new(demo.Showcase), new(*demo.MyShowcase)),
	wire.Struct(new(http.Client)),
	wire.Struct(new(Manager), "*"),
)

type Manager struct {
	SpeechGenerator *demo.SpeechGenerator
	MP3Writer       *demo.MP3Writer
}
