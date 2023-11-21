package jenkins

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func TriggerJenkinsWithoutParam(jenkinsURL string) error {

	username := "yohan"
	token := "115d457e96b15261c7ab907c3628821408"
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + token))
	headers := map[string]string{"Authorization": "Basic " + auth}

	// Construct the POST request
	req, err := http.NewRequest("POST", jenkinsURL, nil)
	if err != nil {
		return err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode > 304 {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}

func TriggerJenkinsWithParams(jenkinsURL string, buildParams map[string]string) error {
	username := "yohan"
	token := "115d457e96b15261c7ab907c3628821408"
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + token))
	headers := map[string]string{
		"Authorization": "Basic " + auth,
		"Content-Type":  "application/x-www-form-urlencoded",
	}

	// Construct the POST request with build parameters
	params := url.Values{}
	for key, value := range buildParams {
		params.Add(key, value)
	}
	requestBody := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", jenkinsURL, requestBody)
	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode > 304 {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}

// Func to provision new folder each time a project initiated

// Func to get list job in certain directory
