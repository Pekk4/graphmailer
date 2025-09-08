package graphmailer

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func HtmlMailer(
	subject string,
	sender string,
	templatePath string,
	attachments []Attachment,
	recipients []string,
	session Session,
	batchSize int,
	batchBreak int,
) {
	htmlContent, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	htmlString := string(htmlContent)
	//attachments := []Attachment{
	//	NewAttachment("logo.png", ExampleLogo, "image/png", "logo.png", "true"),
	//}

	email := NewMessage(
		subject,
		htmlString,
		"HTML",
		sender,
		attachments,
	)

	//for i, address := range recipients {
	for i := 0; i < len(recipients); i++ {
		if i > 0 && i%batchSize == 0 {
			log.Println()

			log.Println(
				batchSize,
				" mails sent, having a ",
				batchBreak,
				" minutes break to avoid throttling the API...",
			)

			time.Sleep(time.Duration(batchBreak) * time.Minute)
			log.Println("Continuing...")
			log.Println()
		}

		address := recipients[i]

		email.Message.To = []To{
			{
				EmailAddress: EmailAddress{
					Address: address,
				},
			},
		}

		jsonData, err := json.Marshal(email)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("POST", session.Endpoint, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+session.Token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.Status != "202 Accepted" {
			log.Println("FATAL - An error occurred - See the response below")
			log.Println("---")
			log.Println(resp)
			log.Println("---")
			log.Println("Going to sleep for 5 minutes and then trying again...")
			time.Sleep(5 * time.Minute)
			log.Println("Okay, let's try again with address: " + address)
			i--
		} else {
			log.Println("Email sent to recipient " + address + " successfully with status: " + resp.Status)
		}
	}

	log.Println("---")
	log.Println("All emails sent successfully, halting...")
}
