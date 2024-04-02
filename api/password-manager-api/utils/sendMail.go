package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"password-manager-service/types"
	"text/template"
)

func SendResetPasswordEmail(email, authToken string) {
	// smtp server configuration.
	emailConfigs := getEmailConfiguration()

	subject := "ResetPassword Link"

	// Authentication.
	auth := smtp.PlainAuth("", emailConfigs.Email, emailConfigs.Password, emailConfigs.SmtpHost)

	var body bytes.Buffer

	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, emailConfigs.MimeHeaders)))

	//Insert required data in email template
	t, _ := template.ParseFiles("templates/template.html")

	t.Execute(&body, struct {
		ResetPasswordLink string
	}{
		ResetPasswordLink: emailConfigs.UIUrl + "/resetPassword?" + "authToken=" + authToken,
	})

	// Sending email.
	err := smtp.SendMail(emailConfigs.SmtpHost+":"+emailConfigs.SmtpPort, auth, emailConfigs.Email, []string{email}, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}

func getEmailConfiguration() types.EmailConfigData {
	var emailConfigs types.EmailConfigData
	emailConfigs.Email = os.Getenv("EMAIL")
	emailConfigs.Password = os.Getenv("EMAIL_KEY")
	emailConfigs.SmtpHost = os.Getenv("SMTP_HOST")
	emailConfigs.SmtpPort = os.Getenv("SMTP_PORT")

	emailConfigs.MimeHeaders = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	emailConfigs.UIUrl = os.Getenv("UI_URL")
	return emailConfigs
}
