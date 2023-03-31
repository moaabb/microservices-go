package data

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadJSON(data io.ReadCloser, v any) error {
	d := json.NewDecoder(data)

	err := d.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, data any, status int) error {
	e := json.NewEncoder(w)

	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")

	err := e.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
