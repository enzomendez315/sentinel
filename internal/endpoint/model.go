package endpoint

import "time"

type Endpoint struct {
	ID        string
	URL       string
	Name      string
	Enabled   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
