package eventstruct

import (
	"net/http"
	"strings"

	"github.com/SpeedVan/go-common-faas/constant/httpconst"
	"github.com/SpeedVan/invoker/service"
)

// RequestEventMetadata todo
type RequestEventMetadata struct {
	ResponseStreamName string                 `json:"responseStreamName"`
	Context            map[string]interface{} `json:"context"`
	Method             string                 `json:"method"`
	Path               string                 `json:"path"`
	Header             http.Header            `json:"header"`
}

// FormHTTPRequest todo
func FormHTTPRequest(r *http.Request, responseStreamName string) *RequestEventMetadata {
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

	return &service.RequestEventMetadata{
		ResponseStreamName: responseStreamName,
		Context:            ctx,
		Method:             r.Method,
		Path:               eventPath,
		Header:             header,
	}
}
