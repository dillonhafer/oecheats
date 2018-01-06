package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func FormatDatabase() {
	dbFile := fmt.Sprintf("%s/cheats-database.xml", openEmuDbLocation)
	cmd := "xmllint"
	args := []string{"--format", dbFile, "--output", dbFile}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		println("Failed to format XML")
		os.Exit(1)
	}

	fmt.Println("Formatted", dbFile)
}

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload local cheat database to source",
	Long:  `WARNING: This will override the main database.xml used by all applications`,
	Run: func(cmd *cobra.Command, args []string) {
		FormatDatabase()
		token := os.Getenv("GITHUB_API_TOKEN")
		if token == "" {
			println("You must provide a GITHUB_API_TOKEN")
			os.Exit(1)
		}

		url := fmt.Sprintf("https://api.github.com/gists/%s", gistId)
		method := "PATCH"

		dbFile := fmt.Sprintf("%s/cheats-database.xml", openEmuDbLocation)
		content, err := ioutil.ReadFile(dbFile)
		if err != nil {
			println(fmt.Sprintf("Could not open %s", dbFile))
			os.Exit(1)
		}

		type Content struct {
			Content string `json:"content"`
		}
		type Files struct {
			File Content `json:"cheats-database.xml"`
		}
		type Gist struct {
			Files Files `json:"files"`
		}

		c := Content{string(content)}
		files := Files{c}
		gist := Gist{files}
		bodyJson, err := json.Marshal(gist)
		if err != nil {
			println("Could not create JSON")
		}

		req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJson))
		req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			fmt.Println("Database updated! ✨")
		} else {
			fmt.Println("Could not update database:", resp.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
