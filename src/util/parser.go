package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseJson(req *http.Request, model interface{}) error {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(req.Body)
	// Read the body of the request
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	// Unmarshal the JSON body into the provided model
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}
