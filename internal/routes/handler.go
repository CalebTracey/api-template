package routes

import (
	"encoding/json"
	request2 "github.com/calebtracey/api-template/external/models/request"
	"github.com/calebtracey/api-template/external/models/response"
	"github.com/calebtracey/api-template/internal/facade"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Handler struct {
	Service facade.APIFacadeI
}

func (h *Handler) InitializeRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Health check
	r.Handle("/health", h.HealthCheck()).Methods(http.MethodGet)

	r.Handle("/add", h.AddNewHandler()).Methods(http.MethodPost)

	staticFs, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(staticFs)
	sh := http.StripPrefix("/swagger-ui/", staticServer)
	r.PathPrefix("/swagger-ui/").Handler(sh)

	return r
}

func (h *Handler) AddNewHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		var psqlResponse response.PSQLResponse
		var psqlRequest request2.PSQLRequest
		defer func() {
			status, _ := strconv.Atoi(psqlResponse.Message.Status)
			hn, _ := os.Hostname()
			psqlResponse.Message.HostName = hn
			psqlResponse.Message.TimeTaken = time.Since(startTime).String()
			_ = json.NewEncoder(writeHeader(w, status)).Encode(psqlResponse)
		}()
		body, bodyErr := readBody(r.Body)

		if bodyErr != nil {
			psqlResponse.Message.ErrorLog = errorLogs([]error{bodyErr}, "Unable to read psqlRequest body", http.StatusBadRequest)
			return
		}
		err := json.Unmarshal(body, &psqlRequest)
		if err != nil {
			psqlResponse.Message.ErrorLog = errorLogs([]error{err}, "Unable to parse psqlRequest", http.StatusBadRequest)
			return
		}

		psqlResponse = h.Service.PSQLResults(r.Context(), psqlRequest)
	}
}

func (h *Handler) HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		if err != nil {
			log.Errorln(err.Error())
			return
		}
	}
}
