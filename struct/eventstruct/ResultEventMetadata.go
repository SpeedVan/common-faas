package eventstruct

import (
	"encoding/json"
	"net/http"
)

// ResultEventMetadata todo
type ResultEventMetadata struct {
	Context       map[string]interface{} `json:"context"`
	ParamsEventID string                 `json:"paramsEventID"`
	StatusCode    int                    `json:"statusCode"`
	Status        string                 `json:"status"`
	Header        http.Header            `json:"header"`
}

// FromHTTPResponseJSONBytes todo
func FromHTTPResponseJSONBytes(bs []byte) (*ResultEventMetadata, error) {
	result := &ResultEventMetadata{}
	err := json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
