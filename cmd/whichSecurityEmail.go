package cmd

import "fmt"

type Topic int

const (
	PHISHING Topic = iota
	WHALING
	RANSOMWARE
	PASSWORD
)

var topics = [...]string{
	"phishing",
	"whaling",
	"ransomware",
	"password",
}

func (topic Topic) String() string {
	return topics[topic-1]
}

// var whalingSubject string = "Cyber Security... What is Whaling?"
var whalingHTML string = "cmd/htmlTemplates/whaling.html"

// var ransomwareSubject string = "Cyber Security is Important team! read this and follow the directions. Thanks all ; )"
var ransomwareHTML string = "cmd/htmlTemplates/ransomware.html"

// var phishingSubject string = "Phishing Awareness. Cyber Security continued..."
var phishingHTML string = "cmd/htmlTemplates/phishing.html"

// var passwordSubject string = "Cyber Security Password tips..."
var passwordHTML string = "cmd/htmlTemplates/password.html"

func WhichSecurityEmail(topic Topic) {
	switch topic {
	case PHISHING:
		fmt.Println("Case Phishing")
		SendSecurityEmail("phishing", phishingHTML)
	case WHALING:
		fmt.Println("Case Whaling")
		SendSecurityEmail("whaling", whalingHTML)
	case RANSOMWARE:
		fmt.Println("Case Ransomware")
		SendSecurityEmail("ransomware", ransomwareHTML)
	case PASSWORD:
		fmt.Println("Case Password")
		SendSecurityEmail("password", passwordHTML)

	}
	fmt.Println("WIP")
}
