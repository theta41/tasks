package ports

import (
	"context"

	"gitlab.com/g6834/team41/tasks/internal/models"
)

type AuthService interface {
	Validate(ctx context.Context, login string, tokens models.TokenPair) (models.TokenPair, error)
}
