package email

import (
	"context"
	"fmt"
	"net/smtp"
)

type EmailSender interface {
	SendPasswordReset(ctx context.Context, toEmail, resetLink string) error
}

type SMTPEmailSender struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func NewSMTPEmailSender(host string, port int, username, password, from string) *SMTPEmailSender {
	return &SMTPEmailSender{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

func (s *SMTPEmailSender) SendPasswordReset(ctx context.Context, toEmail, resetLink string) error {
	subject := "Password Reset Request — staffsearch"
	body := fmt.Sprintf(
		"Hello,\n\nYou requested a password reset. Click the link below to reset your password:\n\n%s\n\nThis link expires in 1 hour.\n\nIf you did not request this, please ignore this email.\n\nstaffsearch Team",
		resetLink,
	)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		s.from, toEmail, subject, body)

	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	return smtp.SendMail(addr, auth, s.from, []string{toEmail}, []byte(msg))
}

// NoOpEmailSender is used when SMTP is not configured (development)
type NoOpEmailSender struct{}

func (n *NoOpEmailSender) SendPasswordReset(ctx context.Context, toEmail, resetLink string) error {
	// In development, just log — no actual email sent
	return nil
}
