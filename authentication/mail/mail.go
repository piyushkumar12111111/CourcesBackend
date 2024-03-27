package mail

import (
    "net/smtp"
)

// EmailConfig holds configuration for the SMTP server
type EmailConfig struct {
    Host     string
    Port     string
    Username string
    Password string
}

// setup SMTP details
var config = EmailConfig{
    Host:     "smtp.example.com",
    Port:     "587",
    Username: "your-username",
    Password: "your-password",
}

// sendMail sends an email
func SendMail(to []string, subject, body string) error {
    auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)

    // The email body
    message := []byte("To: " + to[0] + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" + body)

    // Send the email
    err := smtp.SendMail(config.Host+":"+config.Port, auth, config.Username, to, message)
    return err
}
