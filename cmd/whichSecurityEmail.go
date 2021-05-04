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

var (
	whalingHTML    string = "cmd/htmlTemplates/whaling.html"
	ransomwareHTML string = "cmd/htmlTemplates/ransomware.html"
	phishingHTML   string = "cmd/htmlTemplates/phishing.html"
	passwordHTML   string = "cmd/htmlTemplates/password.html"
)

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
