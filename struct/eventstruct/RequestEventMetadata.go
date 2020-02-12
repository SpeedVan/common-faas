package eventstruct

import "net/http"

// RequestEventMetadata todo
type RequestEventMetadata struct {
	ResponseStreamName string                 `json:"responseStreamName"`
	Context            map[string]interface{} `json:"context"`
	Method             string                 `json:"method"`
	Path               string                 `json:"path"`
	Header             http.Header            `json:"header"`
}
