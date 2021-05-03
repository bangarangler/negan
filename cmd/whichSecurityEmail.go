package cmd

import "fmt"

var Topics = newTopicsInitializer()

func newTopicsInitializer() *topics {
	return &topics{
		Whaling:    "whaling",
		Phishing:   "phishing",
		Ransomware: "ransomware",
		Password:   "password",
	}
}

type topics struct {
	Whaling    string
	Phishing   string
	Ransomware string
	Password   string
}

func WhichSecurityEmail() {
	switch Topics {
	// case Topics.Whaling:
	// 	fmt.Println("Whaling Topic")
	}
	fmt.Println("WIP")
}
