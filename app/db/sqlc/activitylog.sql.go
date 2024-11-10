// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: activitylog.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createActivityLog = `-- name: CreateActivityLog :one
INSERT INTO ActivityLog (petID, activityType, startTime, duration, notes)
VALUES ($1, $2, $3, $4, $5)
RETURNING logid, petid, activitytype, starttime, duration, notes
`

type CreateActivityLogParams struct {
	Petid        pgtype.Int8      `json:"petid"`
	Activitytype string           `json:"activitytype"`
	Starttime    pgtype.Timestamp `json:"starttime"`
	Duration     pgtype.Interval  `json:"duration"`
	Notes        pgtype.Text      `json:"notes"`
}

func (q *Queries) CreateActivityLog(ctx context.Context, arg CreateActivityLogParams) (Activitylog, error) {
	row := q.db.QueryRow(ctx, createActivityLog,
		arg.Petid,
		arg.Activitytype,
		arg.Starttime,
		arg.Duration,
		arg.Notes,
	)
	var i Activitylog
	err := row.Scan(
		&i.Logid,
		&i.Petid,
		&i.Activitytype,
		&i.Starttime,
		&i.Duration,
		&i.Notes,
	)
	return i, err
}

const deleteActivityLog = `-- name: DeleteActivityLog :exec
DELETE FROM ActivityLog WHERE logID = $1
`

func (q *Queries) DeleteActivityLog(ctx context.Context, logid int64) error {
	_, err := q.db.Exec(ctx, deleteActivityLog, logid)
	return err
}

const getActivityLogByID = `-- name: GetActivityLogByID :one
SELECT logid, petid, activitytype, starttime, duration, notes FROM ActivityLog WHERE logID = $1
`

func (q *Queries) GetActivityLogByID(ctx context.Context, logid int64) (Activitylog, error) {
	row := q.db.QueryRow(ctx, getActivityLogByID, logid)
	var i Activitylog
	err := row.Scan(
		&i.Logid,
		&i.Petid,
		&i.Activitytype,
		&i.Starttime,
		&i.Duration,
		&i.Notes,
	)
	return i, err
}

const listActivityLogs = `-- name: ListActivityLogs :many
SELECT logid, petid, activitytype, starttime, duration, notes FROM ActivityLog WHERE petID = $1 ORDER BY startTime DESC LIMIT $2 OFFSET $3
`

type ListActivityLogsParams struct {
	Petid  pgtype.Int8 `json:"petid"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

func (q *Queries) ListActivityLogs(ctx context.Context, arg ListActivityLogsParams) ([]Activitylog, error) {
	rows, err := q.db.Query(ctx, listActivityLogs, arg.Petid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Activitylog{}
	for rows.Next() {
		var i Activitylog
		if err := rows.Scan(
			&i.Logid,
			&i.Petid,
			&i.Activitytype,
			&i.Starttime,
			&i.Duration,
			&i.Notes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateActivityLog = `-- name: UpdateActivityLog :exec
UPDATE ActivityLog
SET activityType = $2, startTime = $3, duration = $4, notes = $5
WHERE logID = $1
`

type UpdateActivityLogParams struct {
	Logid        int64            `json:"logid"`
	Activitytype string           `json:"activitytype"`
	Starttime    pgtype.Timestamp `json:"starttime"`
	Duration     pgtype.Interval  `json:"duration"`
	Notes        pgtype.Text      `json:"notes"`
}

func (q *Queries) UpdateActivityLog(ctx context.Context, arg UpdateActivityLogParams) error {
	_, err := q.db.Exec(ctx, updateActivityLog,
		arg.Logid,
		arg.Activitytype,
		arg.Starttime,
		arg.Duration,
		arg.Notes,
	)
	return err
}