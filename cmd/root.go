/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"

    "github.com/spf13/cobra"
)

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

var messages []Message

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "want",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    // Uncomment the following line if your bare application
    // has an action associated with it:
    // Run: func(cmd *cobra.Command, args []string) { },
    Run: func(cmd *cobra.Command, args []string) {
        endpoint := "https://api.openai.com/v1/chat/completions"
        apiKey := os.Getenv("OPENAI_API_KEY")

        messages = append(messages, Message{
            Role:    "user",
            Content: "who is batman?",
        })

        requestBody, err := json.Marshal(map[string]interface{}{
            "messages": messages,
            "model":    "gpt-3.5-turbo",
        })
        if err != nil {
            fmt.Println(err)
            return
        }

        request, err := http.NewRequest("POST", endpoint, strings.NewReader(string(requestBody)))
        if err != nil {
            fmt.Println(err)
            return
        }
        request.Header.Set("Content-Type", "application/json")
        request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

        client := &http.Client{}
        resp, err := client.Do(request)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Println(err)
            return
        }

        var result map[string]interface{}
        err = json.Unmarshal(body, &result)
        if err != nil {
            fmt.Println(err)
            return
        }

        fmt.Println(result)
    },
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

    // rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.want.yaml)")

    // Cobra also supports local flags, which will only run
    // when this action is called directly.
    rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
