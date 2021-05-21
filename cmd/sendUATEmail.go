package cmd

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

func sendUATEmail(typ string, content string) {
	email := goDotEnvVar("EMAIL_ADDRESS")
	pass := goDotEnvVar("EMAIL_PASSWORD")

	from := email
	password := pass

	f, err := os.Open("./sendToEmails/uatEmails.txt")
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
	if err != nil {
		println(err)
	}

	t, err := template.ParseFiles(content)
	if err != nil {
		println(err)
	}

	var body bytes.Buffer

	senderName := goDotEnvVar("SENDER_NAME")
	// println(typ)

	switch typ {
	case "standard":
		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: DEV -> UAT Updates... \n%s\n\n", mimeHeaders)))
	default:
		fmt.Println("Default case uatEmails")
	}

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    senderName,
		Message: typ,
	})

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("UAT Email Sent!")
}

// ********** LABEL ********** //

var standardHTML string = "cmd/htmlTemplates/standard.html"

// securityEmailsCmd represents the securityEmails command
var uatEmailsCmd = &cobra.Command{
	Use:   "uatEmails",
	Short: "Interactive way to send Nowigence Dev to UAT emails",
	Long: `Email automation from html
	to cli...  write up documentation for UAT in notion.
	export as pdf and move here. fire it off.  Can be removed once
	notion api is up and running.`,
	Run: func(cmd *cobra.Command, args []string) {
		// var req string
		fmt.Println("Entering UAT email section")
		name, _ := cmd.Flags().GetString("emailType")
		// println("name", name)
		switch name {
		case "standard":
			name = "Sending UAT standard Email"
			sendUATEmail("standard", standardHTML)
			// name = "Sending Whaling"
		// 	WhichSecurityEmail(WHALING)
		// case "phishing":
		// 	// name = "Sending Phishing"
		// 	WhichSecurityEmail(PHISHING)
		// case "ransomware":
		// 	// name = "Sending Ransomware"
		// 	WhichSecurityEmail(RANSOMWARE)
		// case "password":
		// 	// name = "Sending Password"
		// 	WhichSecurityEmail(PASSWORD)
		default:
			name = "You must select standard"
		}
		fmt.Println(name)

	},
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	return errors.New("Must use whaling | phishing | ransomware | password")
	// },
}

func init() {
	rootCmd.AddCommand(uatEmailsCmd)
	uatEmailsCmd.Flags().StringVarP(&EmailType, "emailType", "e", "", "Negan will send UAT Email")
}
