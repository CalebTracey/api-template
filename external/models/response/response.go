package response

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

type PSQLResponse struct {
	RowsAffected string  `json:"rowsAffected,omitempty"`
	LastInsertID string  `json:"lastInsertID,omitempty"`
	Message      Message `json:"message,omitempty"`
}

const (
	Success = "SUCCESS"
	Error   = "ERROR"
)

type Message struct {
	ErrorLog  ErrorLogs `json:"errorLog,omitempty"`
	HostName  string    `json:"hostName,omitempty"`
	Status    string    `json:"status,omitempty"`
	TimeTaken string    `json:"timeTaken,omitempty"`
	Count     int       `json:"count,omitempty"`
}

type ErrorLogs []ErrorLog

type ErrorLog struct {
	Scope         string `json:"scope,omitempty"`
	StatusCode    string `json:"status,omitempty"`
	Trace         string `json:"trace,omitempty"`
	RootCause     string `json:"rootCause,omitempty"`
	Query         string `json:"query,omitempty"`
	ExceptionType string `json:"exceptionType,omitempty"`
}

func (m *Message) AddMessageDetails(sw time.Time) {
	host, err := os.Hostname()
	if err != nil {
		log.Error("error retrieving hostname...")
		host = ""
	}

	m.HostName = host
	m.TimeTaken = time.Since(sw).String()
	if len(m.ErrorLog) == 0 {
		m.Status = Success
	} else {
		m.Status = Error
	}
}

func (errs ErrorLogs) GetHTTPStatus(lengthOfResults int) (status int) {
	var s500, s400, s404, s206, s200 bool
	for _, e := range errs {
		code, _ := strconv.Atoi(e.StatusCode)
		switch {
		case code == 206:
			s206 = true
		case code == 404:
			if lengthOfResults > 0 {
				s206 = true
			} else {
				s404 = true
			}
		case code == 400:
			if lengthOfResults > 0 {
				s206 = true
			} else {
				s400 = true
			}
		case code >= 500:
			if lengthOfResults > 0 {
				s206 = true
			} else {
				s500 = true
			}
		default:
			s500 = true
		}
	}
	switch {
	case s206:
		status = http.StatusPartialContent
	case s500:
		status = http.StatusInternalServerError
	case s400:
		status = http.StatusBadRequest
	case s404:
		status = http.StatusNotFound
	case s200:
		status = http.StatusOK
	default:
		status = http.StatusInternalServerError
	}
	return status
}
