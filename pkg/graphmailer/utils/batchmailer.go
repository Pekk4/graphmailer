package graphmailer

// https://learn.microsoft.com/en-us/graph/json-batching
// A sketch of the batch mailer sending just one request instead of n
// But there's a limit of 20 requests per batch, so this is quite useless

//
//import (
//	"log"
//	"os"
//)
//
//type Batch struct {
//	Requests Requests `json:"requests"`
//	// requests should be provided as an array
//}
//
//type Request struct {
//	ID      int     `json:"id"`
//	Method  string  `json:"method"`
//	URL     string  `json:"url"`
//	Body    Email   `json:"body"`
//	Headers Headers `json:"headers"`
//}
//
//type Headers struct {
//	ContentType   string `json:"Content-Type"`
//	Authorization string `json:"Authorization"`
//}
//
//func NewRequest(
//	id int,
//	body Email,
//) Request {
//	request := Request{
//		ID:     id,
//		Method: "POST",
//		URL:    "https://example.com/api/nonnonnoo/",
//		Body:   body,
//		Headers: Headers{
//			ContentType:   "application/json",
//			Authorization: "Bearer ",
//		},
//	}
//
//	return request
//}
//
//func BatchMailer(
//	subject string,
//	sender string,
//	templatePath string,
//	recipients []string,
//	session Session,
//) Batch {
//	htmlContent, err := os.ReadFile(templatePath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	htmlString := string(htmlContent)
//	attachments := []Attachment{}
//
//	email := NewMessage(
//		subject,
//		htmlString,
//		"HTML",
//		sender,
//		attachments,
//	)
//
//	lenkki := len(recipients) + 1
//	requests := make([]Request, lenkki)
//
//	for i, address := range recipients {
//		email.Message.To = []To{
//			{
//				EmailAddress: EmailAddress{
//					Address: address,
//				},
//			},
//		}
//
//		requests[i] = NewRequest(i+1, email)
//	}
//
//	batch := Batch{
//		Requests: requests,
//	}
//	//jsonData, err := json.Marshal(email)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//
//	//req, err := http.NewRequest("POST", session.Endpoint, bytes.NewBuffer(jsonData))
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//req.Header.Set("Content-Type", "application/json")
//	//req.Header.Set("Authorization", "Bearer "+session.Token)
//	//
//	//resp, err := http.DefaultClient.Do(req)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//defer resp.Body.Close()
//	//
//	//if resp.Status != "202 Accepted" {
//	//	log.Println("FATAL - An error occurred - See the response below")
//	//	log.Fatal(resp)
//	//}
//	//
//	//log.Println("Email sent to recipient " + address + " successfully with status: " + resp.Status)
//
//}
//
