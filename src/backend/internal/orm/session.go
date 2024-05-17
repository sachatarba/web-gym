package orm

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ClientID  uuid.UUID
	SessionID uuid.UUID
	TTL       time.Time
}
