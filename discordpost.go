package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getEnvVar(varName string) (result string) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == varName {
			return pair[1]
		}
	}
	return ""
}

// Slack represents ...
type Discord struct {
	Content string `json:"content"`
}

const (
	envWebHookURL = "DISCORD_WEBHOOK_URL"
)

func main() {
	webhookURL := getEnvVar(envWebHookURL)
	if webhookURL == "" {
		fmt.Println(envWebHookURL, "is not specified.")
		os.Exit(1)
	}

	in := os.Stdin
	var buf string
	reader := bufio.NewReaderSize(in, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("failed to read from stdin. err =", err)
			os.Exit(1)
		}
		buf += string(line) + "\n"
	}

	if buf == "" {
		// buf is empty. exit.
		os.Exit(0)
	}

	resp, _ := http.PostForm(
		webhookURL,
		url.Values{"content": []string{buf}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
