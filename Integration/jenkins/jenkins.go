package jenkins

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func TriggerJenkinsWithoutParam() error {
	jenkinsURL := "https://staging-jenkins.nexcloud.id/job/devsync/job/Credential/build"

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
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	fmt.Println("Jenkins job triggered successfully.")
	return nil
}

// Func to provision new folder each time a project initiated

// Func to get list job in certain directory