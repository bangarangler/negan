package cmd

import (
	"bufio"
	"fmt"
	// "io/fs"
	"log"

	"runtime"

	"os"
	"os/exec"
	// "embed"

	// "github.com/kardianos/osext"
	"github.com/pkg/browser"
	// "github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

//go:embedtabs/*
// var e embed.FS

// startdayCmd represents the startday command
var startdayCmd = &cobra.Command{
	Use:   "startday",
	Short: "Open my tabs",
	Long: `Start my day and open
	my common / always open tabs
	that I use for work.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Good morning JP | I'm opening your tabs for work now.")
		browser.OpenURL("https://calendar.google.com/calendar/b/1/r?tab=wc")
		system := runtime.GOOS
		switch system {
		case "windows":
			println("No Thank You, Switch to Linux ; )")
		case "darwin":
			println("Running on mac")
			err := exec.Command("open", "-a", "mailspring").Run()
			err1 := exec.Command("open", "-a", "slack").Run()
			if err != nil || err1 != nil {
				log.Fatal("error", err)
			}
		case "linux":
			println("Linux ; )")
		default:
			fmt.Printf("%s.\n", system)
		}
		// Open the file.
		// this is not relative to this file. it's the path from the root or the
		// path from where it will be ran
		// test, err := fs.Sub(e, "tabs.txt")
		// println("test", test)
		// data, err := e.ReadFile("/tabs/tabs.txt")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// println("data", data)
		f, err := os.Open("./tabs/tabs.txt")

		// myDirFiles, _ := fs.ReadDir(files, "mydir/subdir")
		//     for _, cppFile := range myDirFiles {
		//         fmt.Printf("%q\n", cppFile.Name())
		//     }
		// f, err := os.Open(e)
		if err != nil {
			log.Fatal(err)
		}
		// // Create a new Scanner for the file.
		scanner := bufio.NewScanner(f)
		// // Loop over all lines in the file and print them.
		for scanner.Scan() {
			line := scanner.Text()
			// println(line)
			// use browser to Open the url
			browser.OpenURL(line)
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
