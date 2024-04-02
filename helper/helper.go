package helper

import (
	"fmt"
	"html/template"
	"log"

	"github.com/wneessen/go-mail"
)

// SendMailHTML is a helper function to send an HTML email
func SendMailHTML(subject, templatePath string, to []string, data interface{}) {
	m := mail.NewMsg()
	if err := m.From("roymanlanjutmanik@gmail.com"); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To("roymanlanjutmanik@gmail.com"); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}

	m.Subject(subject)
	m.SetBodyHTMLTemplate(t, data)
	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername("roymanlanjutmanik@gmail.com"), mail.WithPassword("dttuxsvqhoatvuew"))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}
}
