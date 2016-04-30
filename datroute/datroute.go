// Package datroute parses user sent data and combines
// it with the parameters that are passed by the
// [datroute](https://github.com/goaltools/datroute) router.
package datroute

import (
	"log"
	"net/http"
)

// DATRoute is a controller that does two things:
// 1. Calls Request.ParseForm to parse GET / POST requests;
// 2. Makes Request available in your controller (use c.Request).
type DATRoute struct {
	Request *http.Request `bind:"request"`
}

// Before calls ParseForm of the c.Request. Moreover,
// parameters extracted from URN by the datroute router
// are saved to the Form field of the Request.
func (c *DATRoute) Before() http.Handler {
	// Save the old value of Form, datroute router
	// uses it to pass parameters extracted from URN.
	t := c.Request.Form

	// Set r.Form to nil, otherwise ParseForm / ParseMultipartForm will not work.
	c.Request.Form = nil

	// Parse the body depending on the Content-Type.
	var err error
	switch c.Request.Header.Get("Content-Type") {
	default:
		err = c.Request.ParseForm()
	}

	// Make sure the parsing was successfull.
	// Otherwise, return a "bad request" error.
	if err != nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			go log.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		})
	}

	// Add the old values from router to the new r.Form.
	// Copying only one value per key as the router does not pass more than that.
	for k := range t {
		c.Request.Form.Add(k, t.Get(k))
	}
	return nil
}
