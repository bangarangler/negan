package cmd

import (
	// "errors"
	"fmt"

	"github.com/spf13/cobra"
)

var EmailType string

// securityEmailsCmd represents the securityEmails command
var securityEmailsCmd = &cobra.Command{
	Use:   "securityEmails",
	Short: "Interactive way to send Nowigence cyber security emails",
	Long: `I need to send four different emails once a day for 4 days
	every 2 months.  This is automated way to give me options and
	send the correct email to the correct people.`,
	Run: func(cmd *cobra.Command, args []string) {
		// var req string
		fmt.Println("Entering email section")
		name, _ := cmd.Flags().GetString("emailType")
		println("name", name)
		switch name {
		case "whaling":
			name = "Sending Whaling"
		case "phishing":
			name = "Sending Phishing"
		case "ransomware":
			name = "Sending Ransomware"
		case "password":
			name = "Sending Password"
		default:
			name = "You must select whaling | phishing | ransomware | password"
		}
		fmt.Println(name)

	},
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	return errors.New("Must use whaling | phishing | ransomware | password")
	// },
}

func init() {
	rootCmd.AddCommand(securityEmailsCmd)
	securityEmailsCmd.Flags().StringVarP(&EmailType, "emailType", "e", "", "Negan will send Security Email")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// securityEmailsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// securityEmailsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
