/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "io/ioutil"
    "net/http"
    "os"
)

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type OpenaiRequest struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
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
    Run: func(cmd *cobra.Command, args []string) {
        // URL を変数でもつ
        endpoint := "https://api.openai.com/v1/chat/completions"

        // API キー取得
        apiKey := os.Getenv("OPENAI_API_KEY")

        // chatgpt に投げるメッセージを作成
        messages = append(messages, Message{
            Role:    "user",
            Content: "who is batman?",
        })

        requestBody := OpenaiRequest{
            Model:    "gpt-3.5-turbo",
            Messages: messages,
        }

        requestJSON, _ := json.Marshal(requestBody)

        request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestJSON))
        if err != nil {
            fmt.Println(err)
            return
        }
        request.Header.Set("Content-Type", "application/json")
        request.Header.Set("Authorization", "Bearer "+apiKey)

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

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
