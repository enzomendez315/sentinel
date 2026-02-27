-- +goose Up
CREATE TABLE status(
    endpoint_id uuid NOT NULL REFERENCES endpoints(id) ON DELETE CASCADE,
    is_up boolean NOT NULL,
    status_code integer NOT NULL,
    checked_at timestamptz NOT NULL DEFAULT now(),
    latency_ms bigint NOT NULL
    PRIMARY KEY (endpoint_id, checked_at)
);

-- +goose Down
DROP TABLE status;
