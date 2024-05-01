package demo

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
)

// OpenAI defines methods to interact with OpenAI services.
type OpenAI interface {
	GenerateSpeechFromText(
		ctx context.Context,
		body azopenai.SpeechGenerationOptions,
		options *azopenai.GenerateSpeechFromTextOptions,
	) (azopenai.GenerateSpeechFromTextResponse, error)

	GetChatCompletions(
		ctx context.Context,
		body azopenai.ChatCompletionsOptions,
		options *azopenai.GetChatCompletionsOptions,
	) (azopenai.GetChatCompletionsResponse, error)

	GetCompletions(
		ctx context.Context,
		body azopenai.CompletionsOptions,
		options *azopenai.GetCompletionsOptions,
	) (azopenai.GetCompletionsResponse, error)

	GetEmbeddings(
		ctx context.Context,
		body azopenai.EmbeddingsOptions,
		options *azopenai.GetEmbeddingsOptions,
	) (azopenai.GetEmbeddingsResponse, error)

	GetImageGenerations(
		ctx context.Context,
		body azopenai.ImageGenerationOptions,
		options *azopenai.GetImageGenerationsOptions,
	) (azopenai.GetImageGenerationsResponse, error)
}
