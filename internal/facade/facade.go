package facade

import (
	"context"
	"github.com/calebtracey/api-template/external/models"
	"github.com/calebtracey/api-template/external/models/request"
	psql2 "github.com/calebtracey/api-template/internal/facade/psql"
	"strings"
)

type APIFacadeI interface {
	FacadeResponse(ctx context.Context, req models.Request) (resp models.Response)
}

type APIFacade struct {
	PSQLDao psql2.FacadeI
}

func (s APIFacade) FacadeResponse(ctx context.Context, req models.Request) (resp models.Response) {
	//TODO add validation
	if strings.EqualFold(req.Type, "Insert") {
		_ = s.PSQLDao.AddNew(ctx, request.PSQLRequest{})
		//TODO add mappers
	}

	return resp
}
