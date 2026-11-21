package clients

import (
	"context"

	"github.com/google/uuid"
)

type AccessRoleClient interface {
	GetRole(ctx context.Context, userID uuid.UUID) (role string, err error)
}
