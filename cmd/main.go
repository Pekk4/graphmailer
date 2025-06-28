package main

import (
	"github.com/pekk4/graphmailer/pkg/graphmailer"
)

func main() {
	session := graphmailer.InitSession()

	recipients := []string{
		"example@example.com",
	}

	subject := "This is subject!"
	template := "../static/templates/template.html"
	sender := "do-not-reply@example.com"

	graphmailer.HtmlMailer(subject, sender, template, recipients, session, 500, 3)
}
