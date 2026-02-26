package main

import (
	m "github.com/pekk4/graphmailer/pkg/graphmailer"
)

func main() {
	session := m.InitSession()

	recipients := []string{
		"example@example.com",
	}

	subject := "This is subject!"
	template := "../static/templates/template.html"
	sender := "do-not-reply@example.com"
	attachments := []m.Attachment{
		m.NewAttachment("logo.png", m.ExampleLogo, "image/png", "logo.png", "true"),
	}

	m.HtmlMailer(subject, sender, template, attachments, recipients, session, 500, 3)
}
