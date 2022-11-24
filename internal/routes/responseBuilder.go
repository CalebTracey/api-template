package routes

import (
	"bytes"
	"encoding/json"
	"github.com/calebtracey/api-template/external/models/response"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

func writeHeader(w http.ResponseWriter, code int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	return w
}

func renderResponse(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	if _, err = w.Write(content); err != nil {
		logrus.Error(err)
	}
}

func readBody(body io.ReadCloser) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, copyErr := io.Copy(buf, body)

	if copyErr != nil {
		return nil, copyErr
	}
	return buf.Bytes(), nil
}

func errorLogs(errors []error, rootCause string, status int) response.ErrorLogs {
	var errLogs response.ErrorLogs
	for _, err := range errors {
		errLogs = append(errLogs, response.ErrorLog{
			RootCause:  rootCause,
			StatusCode: strconv.Itoa(status),
			Trace:      err.Error(),
		})
	}
	return errLogs
}
