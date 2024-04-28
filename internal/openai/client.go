package openai

import (
	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"

	"github.com/angelokurtis/go-talk/internal/errors"
)

func NewClient(cfg *Config) (*azopenai.Client, error) {
	credential := azcore.NewKeyCredential(cfg.OpenAIKey)
	options := &azopenai.ClientOptions{}

	client, err := azopenai.NewClientForOpenAI(cfg.Endpoint, credential, options)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return client, nil
}
