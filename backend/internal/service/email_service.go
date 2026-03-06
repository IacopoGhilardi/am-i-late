package service

import (
	"github.com/iacopoGhilardi/amILate/internal/config"
	"github.com/iacopoGhilardi/amILate/internal/email"
	log "github.com/iacopoGhilardi/amILate/internal/utils/logger"
	"github.com/resend/resend-go/v2"
)

type EmailService struct {
	client         *resend.Client
	templateEngine *email.TemplateEngine
	cfg            *config.Config
}

func NewEmailService(client *resend.Client, templateEngine *email.TemplateEngine, cfg *config.Config) *EmailService {
	return &EmailService{client: client, templateEngine: templateEngine, cfg: cfg}
}

func (s *EmailService) SendResetPasswordEmail(to, token string) error {
	html, err := s.templateEngine.Render("reset_password", map[string]string{
		"token": token,
		"url":   s.cfg.ResetPasswordURL,
	})

	if err != nil {
		return err
	}

	params := &resend.SendEmailRequest{
		From: s.cfg.EmailFromAddress,
		To:   []string{to},
		Html: html,
	}

	sent, err := s.client.Emails.Send(params)
	if err != nil {
		log.Fatal("failed to send email: %v", err)
	}

	log.Info("email sent: %s", sent.Id)
	return nil
}
