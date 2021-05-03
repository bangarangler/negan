/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// startdayCmd represents the startday command
var startdayCmd = &cobra.Command{
	Use:   "startday",
	Short: "Open my tabs",
	Long: `Start my day and open
	my common / always open tabs
	that I use for work.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Good morning JP | I'm opening your tabs for work now.")
		// Open the file.
		// this is not relative to this file. it's the path from the root or the
		// path from where it will be ran
		f, err := os.Open("./tabs/tabs.txt")
		if err != nil {
			log.Fatal(err)
		}
		// Create a new Scanner for the file.
		scanner := bufio.NewScanner(f)
		// Loop over all lines in the file and print them.
		for scanner.Scan() {
			line := scanner.Text()
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
