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
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Take quick notes",
	Long: `A fast way to take notes
	extend on this in the future hook it up with notion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Opening a new note for you.  -Negan out")
		// create a folder for the notes if doesn't exists
		_ = os.Mkdir("./notes", 0755)
		// the name of the file will be your argument
		// negan notes myfile will open a file called myfile.txt
		editorCmd := exec.Command("nvim", fmt.Sprintf("./notes/%v.txt", args[0]))
		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr

		err := editorCmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
