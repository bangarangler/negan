package cmd

import (
	"bufio"
	"fmt"

	// "io/fs"
	"log"

	"runtime"

	// _ "embed"
	"os"
	"os/exec"

	// "github.com/kardianos/osext"
	"github.com/pkg/browser"
	// "github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

//// go:embed tabs.txt
// var eTest string

// var e []byte

// startdayCmd represents the startday command
var startdayCmd = &cobra.Command{
	Use:   "startday",
	Short: "Open my tabs",
	Long: `Start my day and open
	my common / always open tabs
	that I use for work.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Good morning JP | I'm opening your tabs for work now.")
		system := runtime.GOOS
		switch system {
		case "windows":
			println("No Thank You, Switch to Linux ; )")
		case "darwin":
			println("Running on mac")
			var err error
			err = exec.Command("open", "-a", "/Applications/Firefox Developer Edition.app").Run()
			err = exec.Command("open", "-a", "hey").Run()
			err = exec.Command("open", "-a", "slack").Run()
			err = exec.Command("open", "-a", "/Applications/Microsoft Teams.app").Run()
			if err != nil {
				log.Fatal("error", err)
			}
			f, err := os.Open("/Users/jonathanpalacio/go/tabs/tabs.txt")
			if err != nil {
				log.Fatal(err)
			}
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				browser.OpenURL(line)
			}
		case "linux":
			println("Linux ; )")
			// nohup command >/dev/null 2>&1 &
			err := exec.Command("slack").Start()
			err1 := exec.Command("/usr/local/firefox/firefox-bin").Start() // Firefox Developer Edition
			err2 := exec.Command("hey").Start()
			// err := exec.Command("slack", "/usr/local/firefox/firefox-bin", "mailspring").Start()
			if err != nil || err1 != nil || err2 != nil {
				log.Fatal(err, err1, err2)
			}
			// Open the file.
			// this is not relative to this file. it's the path from the root or the
			// path from where it will be ran
			f, err := os.Open("/home/jonathan/go/tabs/tabs.txt")
			if err != nil {
				log.Fatal(err)
			}
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				browser.OpenURL(line)
			}
		default:
			fmt.Printf("%s.\n", system)
		}
	},
}

func init() {
	rootCmd.AddCommand(startdayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startdayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startdayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
