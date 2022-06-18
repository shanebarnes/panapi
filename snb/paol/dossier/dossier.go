package dossier

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const defaultHostPath = "paol-efel.snb.ca/pas-shim/api/paol/dossier/"

func NewRequest(pan, cookie string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s%s", defaultHostPath, pan), nil)
	if err == nil {
		req.Header.Add("Cookie", cookie)
	}
	return req, err
}

func DoRequest(client *http.Client, req *http.Request) (*Results, error) {
	var results *Results
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			var body []byte
			if body, err = io.ReadAll(resp.Body); err == nil {
				results = &Results{}
				err = json.Unmarshal(body, results)
			}
		} else {
			err = fmt.Errorf("Unexpected response: %v\n", resp)
		}
	}
	return results, err
}
