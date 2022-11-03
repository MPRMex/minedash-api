package minedash

import (
	"encoding/json"
	"io"
	"net/http"
)

func parseJsonBody(res *http.Response, dest interface{}) error {
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &dest)
	if err != nil {
		return err
	}
	return nil
}
