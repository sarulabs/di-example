package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// JSONResponse tries to convert the given data into a json.
// Then it uses the given http.ResponseWriter to writes the json in the response.
// The response headers are updated with 'Content-Type: application/json'
// and the request status is changed to the given status.
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	resp, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// ReadJSONBody reads the body of the given request.
// It tries to decode it, as a json, in the given data interface.
// The data interface must be a pointer.
// If the request body can not be decoded, an error is returned.
func ReadJSONBody(r *http.Request, data interface{}) error {
	body, err := ReadBody(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

// ReadBody reads the body of the given http.Request.
// It closes the request body but it replaces it with a new one.
// So the function can be called multiple times on the same request.
func ReadBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body.Close()

	// replace the body in the request so that it can be red again
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}
