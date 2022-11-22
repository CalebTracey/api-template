package psql

import (
	"database/sql"
	"fmt"
	"github.com/calebtracey/api-template/external/models/request"
	"github.com/calebtracey/api-template/external/models/response"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type MapperI interface {
	CreatePSQLExec(req request.PSQLRequest) string
	PSQLResultToResponse(res sql.Result) (response response.PSQLResponse)
}

type Mapper struct{}

func (m Mapper) CreatePSQLExec(req request.PSQLRequest) string {
	return fmt.Sprintf(exec, req.Table, req.ID)
}

func (m Mapper) PSQLResultToResponse(res sql.Result) (response response.PSQLResponse) {
	if res != nil {
		rowsAffected, rowErr := res.RowsAffected()
		if rowErr != nil {
			log.Errorln(rowErr.Error())
		}
		lastInsertID, idErr := res.LastInsertId()
		if idErr != nil {
			log.Errorln(idErr.Error())
		}
		response.RowsAffected = strconv.Itoa(int(rowsAffected))
		response.LastInsertID = strconv.Itoa(int(lastInsertID))
	}
	return response
}

const exec = `insert into '%s' (id) values ('%s')`
