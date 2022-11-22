package psql

import (
	"context"
	"database/sql"
	"github.com/calebtracey/api-template/external/models/response"
	log "github.com/sirupsen/logrus"
)

type DAOI interface {
	InsertOne(ctx context.Context, exec string) (res sql.Result, error *response.ErrorLog)
}

type DAO struct {
	DB *sql.DB
}

func (s DAO) InsertOne(ctx context.Context, exec string) (res sql.Result, error *response.ErrorLog) {
	res, err := s.DB.ExecContext(ctx, exec)
	if err != nil {
		log.Error(err)
		error = mapError(err, exec)
		return res, error
	}
	return res, nil
}

func mapError(err error, query string) (errLog *response.ErrorLog) {
	errLog = &response.ErrorLog{
		Query: query,
	}
	if err != nil {
		errLog.RootCause = err.Error()
	}
	errLog.StatusCode = "500"
	return errLog
}
