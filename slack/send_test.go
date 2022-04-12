package slack

import (
	"testing"

	"github.com/fldu/unisender/utils"
)

func TestValidateToken(t *testing.T) {
	d := utils.Config{}
	tokens := []string{
		"",
		"test",
	}
	for _, j := range tokens {
		d.Slack.Token = j
		switch {
		case j == "":
			_, err := validateToken(d)
			if err != nil {
				t.Error(err.Error())
			}
		case j == "test":
			_, err := validateToken(d)
			if err != nil {
				t.Log("Regex validation tested.")
			}
		}
	}
}

func TestSendNotification(t *testing.T) {
	c := utils.Config{}
	d := "test"
	err := SendNotification(c, d)
	if err != nil {
		t.Error(err.Error())
	}
	c.Slack.Token = "test"
	err = SendNotification(c, d)
	if err == nil {
		t.Error("Error in token validation")
	}
	c.Slack.Token = "T02T68ABN/B02RDUMHYN7/CPCmOXtIgOOZHDzxxteVrYh4"
	err = SendNotification(c, d)
	if err == nil {
		t.Error("Unexpected error")
	}
}
