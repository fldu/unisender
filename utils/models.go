package utils

type Config struct {
	Slack SlackConfig
	Email EmailConfig
}

type SlackConfig struct {
	Token string
}

type EmailConfig struct {
	From         string
	To           []string
	Subject      string
	Body         string
	SMTPServer   string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}
