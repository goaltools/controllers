package json

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONRenderJSONOK_Error(t *testing.T) {
	// Creating a test handler.
	h := (&JSON{}).RenderJSONOK(map[int]int{})

	// Creating a new server using the handler.
	s1 := httptest.NewServer(h)
	defer s1.Close()

	// Making a request to the server's URL.
	res, _ := http.Get(s1.URL)

	// Validating correctness of the response.
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf(`Incorrect status code. Expected: %d, got: %d.`, http.StatusInternalServerError, res.StatusCode)
	}

	// Making sure the object is rendered correctly.
	exp := http.StatusText(http.StatusInternalServerError) + "\n"
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if string(b) != exp {
		t.Errorf("Incorrect result. Expected: `%s`, got: `%s`.", exp, string(b))
	}
}

func TestJSONRenderJSONOK(t *testing.T) {
	obj := map[string]string{"x": "test object"}

	// Creating a test handler.
	h := (&JSON{}).RenderJSONOK(obj)

	// Creating a new server using the handler.
	s1 := httptest.NewServer(h)
	defer s1.Close()

	// Testing both human readable format and a regular modes.
	for _, v := range []struct {
		exp  string
		prep func()
	}{
		{
			`{"x":"test object"}`, func() {
				*indent = false
			},
		},
		{
			`{
	"x": "test object"
}`, func() {
				*indent = true
			},
		},
	} {
		// Preparing parameters.
		v.prep()

		// Making a request to the server's URL.
		res, _ := http.Get(s1.URL)

		// Validating correctness of the response.
		if res.StatusCode != http.StatusOK {
			t.Errorf(`Incorrect status code. Expected: %d, got: %d.`, http.StatusOK, res.StatusCode)
		}

		// Making sure the object is rendered correctly.
		b, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if string(b) != v.exp {
			t.Errorf("Incorrect result. Expected: `%s`, got: `%s`.", v.exp, string(b))
		}
	}
}
