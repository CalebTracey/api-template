package routes

import (
	"bytes"
	"encoding/json"
	request2 "github.com/calebtracey/api-template/external/models/request"
	"github.com/calebtracey/api-template/external/models/response"
	"github.com/calebtracey/api-template/internal/facade"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
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
			logrus.Errorln(err.Error())
			return
		}
	}
}

func writeHeader(w http.ResponseWriter, code int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	return w
}

func readBody(body io.ReadCloser) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, copyErr := io.Copy(buf, body)

	if copyErr != nil {
		return nil, copyErr
	}
	return buf.Bytes(), nil
}

func errorLogs(errors []error, rootCause string, status int) []response.ErrorLog {
	var errLogs []response.ErrorLog
	for _, err := range errors {
		errLogs = append(errLogs, response.ErrorLog{
			RootCause:  rootCause,
			StatusCode: strconv.Itoa(status),
			Trace:      err.Error(),
		})
	}
	return errLogs
}
