package search

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const defaultHostPath = "paol-efel.snb.ca/pas-shim/api/paol/search"

func NewRequest(street, cookie string) (*http.Request, error) {
	params := url.Values{}
	params.Add("s", street)
	body := strings.NewReader(params.Encode())
	contentLength := strconv.Itoa(body.Len())

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://%s", defaultHostPath), body)
	if err == nil {
		req.Header.Add("Accept", "*/*")
		req.Header.Add("Content-Length", contentLength)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Cookie", cookie)
		req.Header.Add("X-SNB-Assessment-API-Version", "1")
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
