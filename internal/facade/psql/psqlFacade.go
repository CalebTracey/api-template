package psql

import (
	"context"
	"github.com/calebtracey/api-template/external/models/request"
	"github.com/calebtracey/api-template/external/models/response"
	"github.com/calebtracey/api-template/internal/dao/psql"
)

type FacadeI interface {
	AddNew(ctx context.Context, request request.PSQLRequest) response.PSQLResponse
}

type Facade struct {
	PSQL       psql.DAOI
	PSQLMapper psql.MapperI
}

func (s Facade) AddNew(ctx context.Context, request request.PSQLRequest) response.PSQLResponse {
	exec := s.PSQLMapper.CreatePSQLExec(request)

	res, err := s.PSQL.InsertOne(ctx, exec)
	if err != nil {
		return response.PSQLResponse{
			Message: response.Message{
				ErrorLog: []response.ErrorLog{*err},
			},
		}
	}
	result := s.PSQLMapper.PSQLResultToResponse(res)
	return result
}
