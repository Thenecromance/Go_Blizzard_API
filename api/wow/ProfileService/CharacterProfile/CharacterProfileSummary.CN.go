package wow_CharacterProfile

import (
	"context"
	"github.com/Thenecromance/BlizzardAPI/ApiError"
)

func init() {
	CNHookCharacterProfileSummary = func(ctx context.Context, field *CharacterProfileSummaryFields) (any, error) {

		return nil, ApiError.ErrorInDevelopment
	}
}
