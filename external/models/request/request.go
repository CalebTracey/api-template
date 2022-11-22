package request

type PSQLRequest struct {
	RequestType string `json:"requestType,omitempty"`
	Table       string `json:"table,omitempty"`
	ID          string `json:"ID,omitempty"`
}
