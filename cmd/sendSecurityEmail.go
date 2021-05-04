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

func SendSecurityEmail(typ string, content string) {

	email := goDotEnvVar("EMAIL_ADDRESS")
	pass := goDotEnvVar("EMAIL_PASSWORD")
	// senderName := goDotEnvVar("SENDER_NAME")

	// Sender data.
	// subject := "Cyber Security... What is Whaling?"
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

	// t, err := template.ParseFiles("cmd/htmlTemplates/whaling.html")
	t, err := template.ParseFiles(content)
	if err != nil {
		println(err)
	}

	var body bytes.Buffer

	senderName := goDotEnvVar("SENDER_NAME")
	println(typ)
	// subject := typ

	switch typ {
	case "whaling":
		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: Cyber Security... What is Whaling? \n%s\n\n", mimeHeaders)))
	case "phishing":
		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: Phishing Awareness. Cyber Security continued... \n%s\n\n", mimeHeaders)))
	case "ransomware":
		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: Cyber Security is Important team! read this and follow the directions. Thanks all \n%s\n\n", mimeHeaders)))
	case "password":
		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: Cyber Security Password tips... \n%s\n\n", mimeHeaders)))
	default:
		fmt.Println("Something Wrong with typ switch")
	}
	// mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// mimeHeaders := fmt.Sprintf("%s \n\n MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n", typ)
	// body.Write([]byte(fmt.Sprintf("Subject: Cyber Security... What is Whaling? \n%s\n\n", mimeHeaders)))
	// fmt.Println(fmt.Sprintf("%s", typ) + fmt.Sprintf("\n\n%s", mimeHeaders))
	// body.Write([]byte(fmt.Sprintf("%s", typ) + fmt.Sprintf("\n\n%s", mimeHeaders)))
	// body.Write([]byte(mimeHeaders))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		// Name: "Jon Palacio",
		Name: senderName,
		// Message: "Cyber Security... What is Whaling?",
		Message: typ,
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
