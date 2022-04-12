package utils

type Config struct {
	Slack SlackConfig
}

type SlackConfig struct {
	Token string
}
