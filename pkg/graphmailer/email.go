package graphmailer

type Email struct {
	Message Message `json:"message"`
}

type Message struct {
	Subject     string       `json:"subject"`
	Body        Body         `json:"body"`
	To          []To         `json:"toRecipients"`
	From        From         `json:"from"`
	Attachments []Attachment `json:"attachments"`
}

type Body struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

type To struct {
	EmailAddress EmailAddress `json:"emailAddress"`
}

type From struct {
	EmailAddress EmailAddress `json:"emailAddress"`
}

type EmailAddress struct {
	Address string `json:"address"`
}

func NewMessage(
	subject string,
	body string,
	bodyType string,
	sender string,
	attachments []Attachment,
) Email {
	email := Email{
		Message: Message{
			Subject: subject,
			Body: Body{
				ContentType: bodyType,
				Content:     body,
			},
			From: From{
				EmailAddress: EmailAddress{
					Address: sender,
				},
			},
			To: []To{
				{
					EmailAddress: EmailAddress{
						Address: "placeholder@example.com",
					},
				},
			},
			Attachments: attachments,
		},
	}

	return email
}
