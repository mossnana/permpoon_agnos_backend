-- name: CreateTable :exec
CREATE TABLE IF NOT EXISTS logs (
  id UUID NOT NULL DEFAULT gen_random_uuid(),
  endpoint TEXT NOT NULL,
  ip TEXT,
  status_code INTEGER,
  request TEXT NOT NULL,
  response TEXT,
  create_at TIMESTAMP WITH TIME ZONE
);

-- name: CreateLog :one
INSERT INTO logs (
  endpoint, ip, status_code, request, response, create_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;
