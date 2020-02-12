package eventstruct

import "net/http"

// ResponseEventMetadata todo
type ResponseEventMetadata struct {
	Context        map[string]interface{} `json:"context"`
	RequestEventID string                 `json:"requestEventId"`
	StatusCode     int                    `json:"statusCode"`
	Status         string                 `json:"status"`
	Header         http.Header            `json:"header"`
}
