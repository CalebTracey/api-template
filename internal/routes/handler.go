package routes

import (
	"encoding/json"
	"github.com/calebtracey/api-template/external/models"
	"github.com/calebtracey/api-template/internal/facade"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Handler struct {
	Service facade.APIFacadeI
}

func (h *Handler) InitializeRoutes() *gin.Engine {

	r := gin.Default()

	// Health check
	r.Handle(http.MethodGet, "/health", h.HealthCheck())

	r.Handle(http.MethodPost, "/add", h.AddNewHandler())
	//TODO fix this for gin
	//staticFs, err := fs.New()
	//if err != nil {
	//	panic(err)
	//}

	//staticServer := http.FileServer(staticFs)
	//sh := http.StripPrefix("/swagger-ui/", staticServer)

	//r.PathPrefix("/swagger-ui/").Handler(sh)

	return r
}

func (h *Handler) AddNewHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sw := time.Now()
		statusCode := http.StatusOK

		var apiRequest models.Request
		var apiResponse models.Response

		apiResponse = h.Service.FacadeResponse(ctx, *apiRequest.FromJSON(ctx.Request.Body))

		if len(apiResponse.Message.ErrorLog) > 0 {
			statusCode = apiResponse.Message.ErrorLog.GetHTTPStatus(len(apiResponse.Stuff))
		}
		apiResponse.Message.AddMessageDetails(sw)

		ctx.JSON(statusCode, apiResponse)
	}
}

func (h *Handler) HealthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := json.NewEncoder(ctx.Writer).Encode(map[string]bool{"ok": true})
		if err != nil {
			log.Errorln(err.Error())
			return
		}
	}
}
