package json

import (
	"encoding/json"
	"net/http"
)

// Handler is a type that implements http.Handler interface.
type Handler struct {
	obj  interface{}
	code int
}

// Status updates status code that will be returned
// when executing the handler.
func (h Handler) Status(code int) http.Handler {
	h.code = code
	return h
}

// ServeHTTP renders the Handler as JSON object.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var err error

	// Marshal the object either in a human readable format or as is.
	switch *indent {
	case true:
		b, err = json.MarshalIndent(h.obj, *indPrefix, *indString)
	default:
		b, err = json.Marshal(h.obj)
	}

	// Make sure there are no any errors.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the the result to the response.
	w.WriteHeader(h.code)
	w.Write(b)
}
