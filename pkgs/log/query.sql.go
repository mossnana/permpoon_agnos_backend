// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package log

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createLog = `-- name: CreateLog :one
INSERT INTO logs (
  endpoint, ip, status_code, request, response, create_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, endpoint, ip, status_code, request, response, create_at
`

type CreateLogParams struct {
	Endpoint   string
	Ip         pgtype.Text
	StatusCode pgtype.Int4
	Request    string
	Response   pgtype.Text
	CreateAt   pgtype.Timestamptz
}

func (q *Queries) CreateLog(ctx context.Context, arg CreateLogParams) (Log, error) {
	row := q.db.QueryRow(ctx, createLog,
		arg.Endpoint,
		arg.Ip,
		arg.StatusCode,
		arg.Request,
		arg.Response,
		arg.CreateAt,
	)
	var i Log
	err := row.Scan(
		&i.ID,
		&i.Endpoint,
		&i.Ip,
		&i.StatusCode,
		&i.Request,
		&i.Response,
		&i.CreateAt,
	)
	return i, err
}

const createTable = `-- name: CreateTable :exec
CREATE TABLE IF NOT EXISTS logs (
  id UUID NOT NULL DEFAULT gen_random_uuid(),
  endpoint TEXT NOT NULL,
  ip TEXT,
  status_code INTEGER,
  request TEXT NOT NULL,
  response TEXT,
  create_at TIMESTAMP WITH TIME ZONE
)
`

func (q *Queries) CreateTable(ctx context.Context) error {
	_, err := q.db.Exec(ctx, createTable)
	return err
}
