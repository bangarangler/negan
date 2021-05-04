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

func WhichSecurityEmail(topic Topic) {
	switch topic {
	case PHISHING:
		fmt.Println("Case Phishing")
	case WHALING:
		fmt.Println("Case Whaling")
		SendWhalingEmail()
	case RANSOMWARE:
		fmt.Println("Case Ransomware")
	case PASSWORD:
		fmt.Println("Case Password")

	}
	fmt.Println("WIP")
}
