package types

import "github.com/google/uuid"

type AuthenticatedUser struct {
	UUID          uuid.UUID
	Username      string
	Email         string
	Authenticated bool
}

// AlertEnumType is a type of alert
type AlertEnumType string

var AlertEnum = struct {
	Success AlertEnumType
	Info    AlertEnumType
	Warning AlertEnumType
	Error   AlertEnumType
}{
	Success: "success",
	Info:    "info",
	Warning: "warning",
	Error:   "error",
}

type AlertType struct {
	Type    AlertEnumType
	Message string
}

type ContextEnumKey string

var ContextEnum = struct {
	User ContextEnumKey
}{
	User: "user",
}
