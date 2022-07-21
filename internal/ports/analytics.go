package ports

import (
	"context"
)

type AnalyticsService interface {
	CreateTask(ctx context.Context, objectId uint32) error
	FinishTask(ctx context.Context, objectId uint32) error
	CreateLetter(ctx context.Context, objectId uint32, email string) error
	AcceptedLetter(ctx context.Context, objectId uint32, email string) error
	DeclinedLetter(ctx context.Context, objectId uint32, email string) error
}
