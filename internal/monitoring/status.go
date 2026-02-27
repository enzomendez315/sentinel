package monitoring

import "time"

type Status struct {
	EndpointID string    `json:"endpoint_id"`
	IsUp       bool      `json:"is_up"`
	StatusCode int       `json:"status_code"`
	CheckedAt  time.Time `json:"checked_at"`
	LatencyMs  int64     `json:"latency_ms"`
}
