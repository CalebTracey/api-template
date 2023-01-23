package models

import (
	"encoding/json"
	"github.com/calebtracey/api-template/external/models/response"
	log "github.com/sirupsen/logrus"
	"io"
)

type Request struct {
	FieldOne string `json:"field_one"`
	Type     string `json:"type"`
}

func (r *Request) FromJSON(reader io.Reader) *Request {
	if err := json.NewDecoder(reader).Decode(r); err != nil {
		log.Errorf("error decoding leaderboard request: %v", err)
	}
	return r
}

type Response struct {
	Stuff   []struct{}
	Message response.Message `json:"message"`
}
