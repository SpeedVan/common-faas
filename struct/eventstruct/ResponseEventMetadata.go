package eventstruct

import (
	"encoding/json"
	"net/http"
)

// ResponseEventMetadata todo
type ResponseEventMetadata struct {
	Context        map[string]interface{} `json:"context"`
	RequestEventID string                 `json:"requestEventId"`
	StatusCode     int                    `json:"statusCode"`
	Status         string                 `json:"status"`
	Header         http.Header            `json:"header"`
}

// FromHTTPResponseJSONBytes todo
func FromHTTPResponseJSONBytes(bs []byte) (*ResponseEventMetadata, error) {
	result := &ResponseEventMetadata{}
	err := json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
