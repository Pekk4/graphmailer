package graphmailer

// DummyMailer is a kind of mock function to simulate sending emails
// TODO: implement tests and move this functionality to there

//
//import (
//	//"bytes"
//	"encoding/json"
//	"log"
//
//	//"net/http"
//	"os"
//	"time"
//)
//
//type Resp struct {
//	status string
//}
//
//var calls []int
//var vipu bool
//
//func HtmlMailer(
//	subject string,
//	sender string,
//	templatePath string,
//	recipients []string,
//	session Session,
//	batchSize int,
//	batchBreak int,
//) {
//	htmlContent, err := os.ReadFile(templatePath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	htmlString := string(htmlContent)
//	attachments := []Attachment{} // OBS! removed logo here
//
//	email := NewMessage(
//		subject,
//		htmlString,
//		"HTML",
//		sender,
//		attachments,
//	)
//
//	//for i, address := range recipients {
//	for i := 0; i < len(recipients); i++ {
//		if i > 0 && i%batchSize == 0 {
//			log.Println()
//
//			log.Println(
//				batchSize,
//				" mails sent, having a ",
//				batchBreak,
//				" minutes break to avoid throttling the API...",
//			)
//
//			time.Sleep(time.Duration(batchBreak) * time.Minute)
//			log.Println("Continuing...")
//			log.Println()
//		}
//
//		address := recipients[i]
//
//		email.Message.To = []To{
//			{
//				EmailAddress: EmailAddress{
//					Address: address,
//				},
//			},
//		}
//
//		jsonData, err := json.Marshal(email)
//		if err != nil {
//			log.Fatal(err)
//		}
//		resp := fakeMail(jsonData)
//		//req, err := http.NewRequest("POST", session.Endpoint, bytes.NewBuffer(jsonData))
//		//if err != nil {
//		//	log.Fatal(err)
//		//}
//		//req.Header.Set("Content-Type", "application/json")
//		//req.Header.Set("Authorization", "Bearer "+session.Token)
//		//
//		//resp, err := http.DefaultClient.Do(req)
//		//if err != nil {
//		//	log.Fatal(err)
//		//}
//		//defer resp.Body.Close()
//
//		//if resp.Status != "202 Accepted" {
//		if resp.status != "202 Accepted" {
//			log.Println("FATAL - An error occurred - See the response below")
//			//log.Fatal(resp)
//			log.Println("---")
//			log.Println(resp)
//			log.Println("---")
//			log.Println("Going to sleep for 10 minutes and then trying again...")
//			time.Sleep(10 * time.Second)
//			log.Println("Okay, let's try again with address: " + address)
//			i--
//		} else {
//			log.Println("Email sent to recipient " + address + " successfully with status: " + resp.status)
//		}
//
//		//log.Println("Email sent to recipient " + address + " successfully with status: " + resp.Status)
//	}
//
//	log.Println("---")
//	log.Println("All emails sent successfully, halting...")
//}
//
//func fakeMail(data []byte) Resp {
//	calls = append(calls, len(calls)+1)
//
//	if len(calls) == 1 && !vipu {
//		log.Println("Ei tuu")
//		vipu = true
//		return Resp{status: "500 Internal Server Error"}
//	}
//
//	return Resp{status: "202 Accepted"}
//}
//
