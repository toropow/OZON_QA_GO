package client_db

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/act-device-api/internal/database"
	"github.com/ozonmp/act-device-api/test/internal/models"

	_ "github.com/lib/pq"
)

type EventStorage interface {
}

type storage struct {
	DB *sqlx.DB
}

func NewPostgres(dsn string) (*storage, error) {
	db, err := database.NewPostgres(dsn, "postgres")
	if err != nil {
		return nil, err
	}
	return &storage{DB: db}, nil
}

func (r storage) ByDeviceId(ctx context.Context, deviceID int) (*models.DeviceEvent, error) {
	var (
		event models.DeviceEvent
	)
	query := sq.Select("id", "device_id", "type", "status", "payload", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From("devices_events").
		Where(sq.Eq{"device_id": deviceID})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	s, args, err = query.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.DB.GetContext(ctx, &event, s, args...)

	return &event, err
}
