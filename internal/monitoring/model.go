package monitoring

import "time"

type Status struct {
	EndpointID string
	IsUp       bool
	StatusCode int
	CheckedAt  time.Time
	LatencyMs  int64
}
