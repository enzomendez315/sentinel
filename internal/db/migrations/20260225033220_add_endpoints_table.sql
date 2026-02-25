-- +goose Up
CREATE TABLE endpoints(
    id uuid NOT NULL PRIMARY KEY,
    url text NOT NULL,
    name text NOT NULL,
    enabled boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE endpoints;
