// Package json provides functions for rendering of JSON objects.
// It uses a standard "encoding/json" package that relies on
// reflection package.
package json

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

var (
	indent = flag.Bool("json:indent", false, "use a human readable format for JSON rendering")

	indPrefix = flag.String("json:indent.prefix", "", "JSON prefix of an indented line")
	indString = flag.String("json:indent.string", "\t", "JSON indentation symbol")
)

// JSON is a controller that provides JSON rendering
// functionality.
type JSON struct {
}

// RenderJSON gets any object and a status code to use
// as input and returns a handler function that can be used
// for rendering of the result.
func (c *JSON) RenderJSON(obj interface{}, status int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var b []byte
		var err error

		// Marshal the object either in a human readable format or as is.
		switch *indent {
		case true:
			b, err = json.MarshalIndent(obj, *indPrefix, *indString)
		default:
			b, err = json.Marshal(obj)
		}

		// Make sure there are no any errors.
		if err != nil {
			go log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Write the the result to the response.
		w.WriteHeader(status)
		w.Write(b)
	})
}

// RenderJSONOK is an equivalent of calling RenderJSON(obj, http.StatusOK).
func (c *JSON) RenderJSONOK(obj interface{}) http.Handler {
	return c.RenderJSON(obj, http.StatusOK)
}
