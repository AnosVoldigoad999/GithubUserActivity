/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gat [username]",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please wait...")
		GetActivity(args[0])
	},
}

func GetActivity(username string) {
	var items []Activity
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &items)
		for _, item := range items {
			PrintOutActivity(item)
		}
		Newbody, _ := json.MarshalIndent(items, "", " ")
		os.WriteFile("github.json", Newbody, 0644)
		fmt.Println("Done")
	} else {
		fmt.Println("An error occurred while calling the api:", err)
	}

	defer resp.Body.Close()
}

func PrintOutActivity(activity Activity) {
	var numOfCommits int
	eventType := activity.Type

	switch eventType {
	case "CommitCommentEvent":
		fmt.Printf("--Created a commit comment in %s \n", activity.Repo.Name)
	case "CreateEvent":
		fmt.Printf("--Created a new branch in %s \n", activity.Repo.Name)
	case "DeleteEvent":
		fmt.Printf("--Deleted a branch in %s \n", activity.Repo.Name)
	case "DiscussionEvent":
		fmt.Printf("--Discussion created in %s \n", activity.Repo.Name)
	case "ForkEvent":
		fmt.Printf("--Forked %s \n", activity.Repo.Name)
	case "GollumEvent":
		fmt.Printf("--%s in %s \n", activity.Payload.Action, activity.Repo.Name)
	case "IssueCommentEvent":
		fmt.Printf("--%s in %s \n", activity.Payload.Action, activity.Repo.Name)
	case "IssuesEvent":
		fmt.Printf("--%s in %s \n", activity.Payload.Action, activity.Repo.Name)
	case "MemberEvent":
		fmt.Printf("--%s in %s \n", activity.Payload.Action, activity.Repo.Name)
	case "PublicEvent":
		fmt.Printf("--Made %s Public \n", activity.Repo.Name)
	case "PullRequestEvent":
		fmt.Printf("--%s in %s \n", activity.Payload.Action, activity.Repo.Name)
	case "PullRequestReviewEvent":
		fmt.Printf("--%s in %s \n", activity.Payload.Action, activity.Repo.Name)
	case "PullRequestReviewCommentEvent":
		fmt.Printf("--%s in %s \n", activity.Payload.Action, activity.Repo.Name)
	case "PushEvent":
		numOfCommits = activity.Payload.Size
		if numOfCommits > 1 || numOfCommits == 0 {
			fmt.Printf("--Pushed %d commits to %s \n", numOfCommits, activity.Repo.Name)
		} else {
			fmt.Printf("--Pushed %d commit to %s \n", numOfCommits, activity.Repo.Name)
		}
	case "ReleaseEvent":
		fmt.Printf("--%s in %s \n", activity.Payload.Action, activity.Repo.Name)
	case "WatchEvent":
		fmt.Printf("--Starred %s \n", activity.Repo.Name)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gac.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
