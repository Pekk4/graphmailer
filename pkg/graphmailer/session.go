package graphmailer

import (
	"fmt"
	"log"
	"os"
)

type Session struct {
	Token    string
	Endpoint string
}

func InitSession() Session {
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	tenantId := os.Getenv("TENANT_ID")
	userId := os.Getenv("USER_ID")
	endpoint := fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/sendMail", userId)

	switch {
	case clientId == "":
		log.Fatal("No CLIENT_ID defined!")
	case clientSecret == "":
		log.Fatal("No CLIENT_SECRET defined!")
	case tenantId == "":
		log.Fatal("No TENANT_ID defined!")
	case userId == "":
		log.Fatal("No USER_ID defined!")
	}

	token := GetAccessToken(clientId, clientSecret, tenantId)

	session := Session{
		Token:    token,
		Endpoint: endpoint,
	}

	return session
}
