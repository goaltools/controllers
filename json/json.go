// Package json provides functions for rendering of JSON objects.
// It uses a standard "encoding/json" package that relies on
// reflection package.
package json

import (
	"flag"
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

// RenderJSON gets any object as input and returns a type
// that implements http.Handler interface.
func (c *JSON) RenderJSON(obj interface{}) http.Handler {
	return Handler{obj, http.StatusOK}
}
