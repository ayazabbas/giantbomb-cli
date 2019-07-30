package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/ayazabbas/giantbombcli/internal/model"
)

// ConvertToMap converts a generic interface to a string map.
func ConvertToMap(params interface{}) map[string]string {
	var i map[string]string

	if params == nil {
		return i
	}

	j, _ := json.Marshal(params)

	err := json.Unmarshal([]byte(j), &i)

	if err != nil {
		fmt.Println(err)
	}

	return i
}

// NewRequest assembles a new request to be executed.
func NewRequest(method string, path string, params interface{}) (*http.Request, error) {
	baseURL, err := url.Parse(model.BaseURL)
	if err != nil {
		return nil, err
	}

	u, _ := baseURL.Parse(path)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()

	q.Add("api_key", model.APIKey)
	q.Add("format", "json")

	p := ConvertToMap(params)

	for key, value := range p {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func DoRequest(req *http.Request, v interface{}) error {
	// Create client to perform request
	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	decErr := json.Unmarshal(buff, &v)

	if decErr == io.EOF {
		decErr = nil
	}
	if decErr != nil {
		err = decErr
	}

	return err
}
