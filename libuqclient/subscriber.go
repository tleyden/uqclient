package uqclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type UniqushSubscriber struct {
	UniqushService UniqushService
	DeviceToken    string
	Name           string
}

// Creates the subscriber on the Uniqush server
func (s UniqushSubscriber) Create() (response string, err error) {

	endpointUrl := fmt.Sprintf("%s/subscribe", s.UniqushService.UniqushClient.UniqushURL)

	formValues := url.Values{
		"service":         {s.UniqushService.Name},
		"subscriber":      {s.Name},
		"pushservicetype": {fmt.Sprintf("%s", s.UniqushService.PushServiceType)},
		"devtoken":        {s.DeviceToken},
	}

	resp, err := http.PostForm(endpointUrl, formValues)
	if err != nil {
		return "", fmt.Errorf("Error in form POST creating subscriber: %+v. Error: %v", formValues, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body creating subscriber: %+v. Error: %v", formValues, err)
	}

	return string(body), nil

}

// Sends a push notificatin via the Uniqush server
func (s UniqushSubscriber) Push(message string) (response string, err error) {

	endpointUrl := fmt.Sprintf("%s/push", s.UniqushService.UniqushClient.UniqushURL)
	formValues := url.Values{
		"service":    {s.UniqushService.Name},
		"subscriber": {s.Name},
		"msg":        {message},
	}

	resp, err := http.PostForm(endpointUrl, formValues)
	if err != nil {
		return "", fmt.Errorf("Error in form POST sending push to: %+v. Error: %v", s, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reaading response body sending push to: %+v. Error: %v", s, err)
	}

	return string(body), nil

}
