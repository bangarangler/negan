package cmd

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func goDotEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}

func SendWhalingEmail() {

	email := goDotEnvVar("EMAIL_ADDRESS")
	pass := goDotEnvVar("EMAIL_PASSWORD")
	senderName := goDotEnvVar("SENDER_NAME")

	// Sender data.
	subject := "Cyber Security... What is Whaling?"
	from := email
	password := pass

	// Receiver email address.
	f, err := os.Open("./sendToEmails/emails.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	to := lines

	// smtp server configuration.
	smtpHost := "smtp.office365.com"
	smtpPort := "587"

	conn, err := net.Dial("tcp", "smtp.office365.com:587")
	if err != nil {
		println(err)
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		println(err)
	}

	tlsconfig := &tls.Config{
		ServerName: smtpHost,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		println(err)
	}

	auth := LoginAuth(from, password)

	if err = c.Auth(auth); err != nil {
		println(err)
	}

	t, err := template.ParseFiles("cmd/htmlTemplates/whaling.html")
	if err != nil {
		println(err)
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// body.Write([]byte(fmt.Sprintf("Subject: Nowigence Cyber Security... \n%s\n\n", mimeHeaders)))
	fmt.Println(fmt.Sprintf("%v\n%s\n\n", subject, mimeHeaders))
	body.Write([]byte(fmt.Sprintf("%s\n%s\n\n", subject, mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		// Name:    "Jon Palacio",
		Name: senderName,
		// Message: "Nowigence Cyper Security...",
		Message: subject,
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
