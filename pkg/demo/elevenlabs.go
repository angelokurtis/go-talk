package demo

import (
	"encoding/json"
	"net/http"

	"github.com/angelokurtis/go-talk/internal/errors"
)

type ElevenLabsAPI struct {
	client *http.Client
}

func NewElevenLabs(client *http.Client) *ElevenLabsAPI {
	return &ElevenLabsAPI{client: client}
}

func (api *ElevenLabsAPI) GetVoices() ([]*Voice, error) {
	request, err := http.NewRequest("GET", "https://api.elevenlabs.io/v1/voices", nil)
	if err != nil {
		return nil, errors.Errorf("failed to create request for voices: %w", err)
	}

	response, err := api.client.Do(request)
	if err != nil {
		return nil, errors.Errorf("request for voices failed: %w", err)
	}

	var voices *VoicesResponse
	if err = json.NewDecoder(response.Body).Decode(&voices); err != nil {
		return nil, errors.Errorf("failed to decode voices response: %w", err)
	}

	if voices == nil {
		return nil, errors.New("decoded voices response is nil")
	}

	return voices.Voices, nil
}

type VoicesResponse struct {
	Voices []*Voice `json:"voices,omitempty"`
}

type Voice struct {
	VoiceID                 *string                   `json:"voice_id,omitempty"`
	Name                    *string                   `json:"name,omitempty"`
	Samples                 any                       `json:"samples"`
	Category                *Category                 `json:"category,omitempty"`
	FineTuning              *FineTuning               `json:"fine_tuning,omitempty"`
	Labels                  *Labels                   `json:"labels,omitempty"`
	Description             any                       `json:"description"`
	PreviewURL              *string                   `json:"preview_url,omitempty"`
	AvailableForTiers       []any                     `json:"available_for_tiers,omitempty"`
	Settings                any                       `json:"settings"`
	Sharing                 any                       `json:"sharing"`
	HighQualityBaseModelIDs []*HighQualityBaseModelID `json:"high_quality_base_model_ids,omitempty"`
	SafetyControl           any                       `json:"safety_control"`
	VoiceVerification       *VoiceVerification        `json:"voice_verification,omitempty"`
	OwnerID                 any                       `json:"owner_id"`
	PermissionOnResource    any                       `json:"permission_on_resource"`
}

type FineTuning struct {
	IsAllowedToFineTune         *bool               `json:"is_allowed_to_fine_tune,omitempty"`
	FinetuningState             *FinetuningState    `json:"finetuning_state,omitempty"`
	VerificationFailures        []any               `json:"verification_failures,omitempty"`
	VerificationAttemptsCount   *int64              `json:"verification_attempts_count,omitempty"`
	ManualVerificationRequested *bool               `json:"manual_verification_requested,omitempty"`
	Language                    *string             `json:"language"`
	FinetuningProgress          *FinetuningProgress `json:"finetuning_progress,omitempty"`
	Message                     any                 `json:"message"`
	DatasetDurationSeconds      any                 `json:"dataset_duration_seconds"`
	VerificationAttempts        any                 `json:"verification_attempts"`
	SliceIDS                    any                 `json:"slice_ids"`
	ManualVerification          any                 `json:"manual_verification"`
}

type FinetuningProgress struct{}

type Labels struct {
	Accent            *string `json:"accent,omitempty"`
	Description       *string `json:"description,omitempty"`
	Age               *Age    `json:"age,omitempty"`
	Gender            *Gender `json:"gender,omitempty"`
	UseCase           *string `json:"use case,omitempty"`
	LabelsDescription *string `json:"description ,omitempty"`
	Featured          *string `json:"featured,omitempty"`
	Usecase           *string `json:"usecase,omitempty"`
}

type VoiceVerification struct {
	RequiresVerification      *bool  `json:"requires_verification,omitempty"`
	IsVerified                *bool  `json:"is_verified,omitempty"`
	VerificationFailures      []any  `json:"verification_failures,omitempty"`
	VerificationAttemptsCount *int64 `json:"verification_attempts_count,omitempty"`
	Language                  any    `json:"language"`
	VerificationAttempts      any    `json:"verification_attempts"`
}

type Category string

const (
	Premade Category = "premade"
)

type FinetuningState string

const (
	FineTuned  FinetuningState = "fine_tuned"
	NotStarted FinetuningState = "not_started"
)

type HighQualityBaseModelID string

const (
	ElevenMultilingualV1 HighQualityBaseModelID = "eleven_multilingual_v1"
	ElevenTurboV2        HighQualityBaseModelID = "eleven_turbo_v2"
)

type Age string

const (
	MiddleAged Age = "middle aged"
	Old        Age = "old"
	Young      Age = "young"
)

type Gender string

const (
	Female Gender = "female"
	Male   Gender = "male"
)
