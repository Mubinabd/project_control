package email

import (
	"net/smtp"

	"github.com/Mubinabd/project_control/internal/genproto/auth"
	"golang.org/x/exp/slog"
)

func SendVerificationCode(params *auth.Params) error {
	msg := []byte("To: " + params.To + "\r\n" +
		"Subject: Email Verification\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		"<html>" +
		"<body>" +
		"<h1>Hi,</h1>" +
		"<p></p>" +
		"<p>Kod: <strong>" + params.Code + "</strong></p>" +
		"</body>" +
		"</html>\r\n")

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("smtp", params.From, params.Password, "smtp.gmail.com"),
		params.From, []string{params.To}, []byte(msg),
	)

	if err != nil {
		slog.Error("Could not send an email", err.Error())
		return err
	}

	return nil
}