package procroute

import (
	"errors"
	"net/http"
)

var (
	ErrHttpResponseWriterNotSet = errors.New("http response writer is not set")
)

// HttpError represents the datatype used as error response
type HttpError struct {
	Status    int
	ErrorCode string
	Message   string
}

// write marshals the error message and sends it back to the client
func (h *HttpError) write(contentType string, parser Parser, w http.ResponseWriter) error {
	if w == nil {
		return ErrHttpResponseWriterNotSet
	}

	bts, err := parser.Marshal(h)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", contentType)
	w.WriteHeader(h.Status)
	w.Write(bts)
	return nil
}
