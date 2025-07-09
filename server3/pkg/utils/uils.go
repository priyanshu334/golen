package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads and unmarshals JSON body into the provided interface.
// It returns an error if anything goes wrong.
func ParseBody(r *http.Request, out interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, out); err != nil {
		return err
	}
	return nil
}
