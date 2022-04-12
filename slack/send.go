package slack

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/fldu/unisender/utils"
)

func SendNotification(c utils.Config, payload string) error {
	ok, err := validateToken(c)
	if err != nil {
		return errors.New(err.Error())
	}
	if !ok {
		return nil
	}
	url := "https://hooks.slack.com/services/" + c.Slack.Token
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return errors.New("error while building HTTP request for Slack notification: " + err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	r, err := client.Do(req)
	if err != nil {
		return errors.New("error while sending notification to Slack: " + err.Error())
	}
	defer client.CloseIdleConnections()
	if r.StatusCode >= 400 {
		err, _ := ioutil.ReadAll(r.Body)
		errData := string(err)
		return errors.New("something went wrong while sending notification to Slack:\n" + errData)
	}
	return nil
}

func validateToken(c utils.Config) (bool, error) {
	r := regexp.MustCompile(`^[A-Z0-9]{9}/[A-Z0-9]{11}/[a-zA-Z0-9]{24}$`)
	switch {
	case r.MatchString(c.Slack.Token):
		return true, nil
	case c.Slack.Token == "":
		return false, nil
	default:
		err := errors.New("slack token is not valid")
		return false, err
	}
}
