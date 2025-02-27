// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package log

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Log struct {
	ID         pgtype.UUID
	Endpoint   string
	Ip         pgtype.Text
	StatusCode pgtype.Int4
	Request    string
	Response   pgtype.Text
	CreateAt   pgtype.Timestamptz
}
