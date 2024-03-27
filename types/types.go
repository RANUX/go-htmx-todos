package types

import "github.com/google/uuid"

type AuthenticatedUser struct {
	UUID          uuid.UUID
	Username      string
	Email         string
	Authenticated bool
}
