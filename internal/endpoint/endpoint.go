package endpoint

import (
	"time"

	"github.com/google/uuid"
)

type Endpoint struct {
	ID        uuid.UUID `json:"id"`
	URL       string    `json:"url"`
	Name      string    `json:"name"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
