package demo

import (
	"context"
	"io"
	"log/slog"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/gotidy/ptr"

	"github.com/angelokurtis/go-talk/internal/errors"
)

type SpeechGenerator struct {
	openAI OpenAI
}

func NewSpeechGenerator(openAI OpenAI) *SpeechGenerator {
	return &SpeechGenerator{openAI: openAI}
}

func (sg *SpeechGenerator) GenerateSpeech(ctx context.Context, input string) (io.ReadCloser, error) {
	body := azopenai.SpeechGenerationOptions{
		DeploymentName: ptr.Of("tts-1"),
		Voice:          ptr.Of(azopenai.SpeechVoiceAlloy),
		ResponseFormat: ptr.Of(azopenai.SpeechGenerationResponseFormatMp3),
		Input:          ptr.Of(input),
	}

	slog.DebugContext(ctx, "Generating speech from text",
		slog.String("input", *body.Input),
		slog.String("voice", string(*body.Voice)),
		slog.String("format", string(*body.ResponseFormat)),
		slog.String("deployment", *body.DeploymentName),
	)

	res, err := sg.openAI.GenerateSpeechFromText(ctx, body, &azopenai.GenerateSpeechFromTextOptions{})
	if err != nil {
		return nil, errors.Errorf("failed to generate speech: %w", err)
	}

	slog.InfoContext(ctx, "Speech generated")

	return res.Body, nil
}
