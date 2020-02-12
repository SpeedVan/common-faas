package eventstruct

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SpeedVan/go-common-faas/constant/httpconst"
)

// RequestEventMetadata todo
type RequestEventMetadata struct {
	ResponseStreamName string                 `json:"responseStreamName"`
	Context            map[string]interface{} `json:"context"`
	Method             string                 `json:"method"`
	Path               string                 `json:"path"`
	Header             http.Header            `json:"header"`
}

// FromHTTPRequestJSONBytes todo
func FromHTTPRequestJSONBytes(bs []byte) (*RequestEventMetadata, error) {
	result := &RequestEventMetadata{}
	err := json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FromHTTPRequest todo
func FromHTTPRequest(r *http.Request, responseStreamName, requestEventPath string) *RequestEventMetadata {
	header := r.Header
	ctx := make(map[string]interface{})

	for k, item := range header {
		if strings.HasPrefix(k, httpconst.HeaderPrefix) {
			if len(item) > 0 {
				ctx[k] = item[0]
			} else {
				ctx[k] = ""
			}
			header.Del(k)
		}
	}

	return &RequestEventMetadata{
		ResponseStreamName: responseStreamName,
		Context:            ctx,
		Method:             r.Method,
		Path:               requestEventPath,
		Header:             header,
	}
}
