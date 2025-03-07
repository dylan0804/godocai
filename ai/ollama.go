package ai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dylan0804/godocai/shared"
)

type Request struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool `json:"stream"`
}

type Response struct {
	Response string `json:"response"`
	Done bool `json:"done"`
}

type AIStreamMsg struct {
	Content string
	Done    bool
	Error   error
}

func Generate(prompt string, updateFn func(string, bool)) error {
	request := Request{
		Model: "deepseek-r1:14b",
		Prompt: fmt.Sprintf(`Generate a practical explanation for the Go package '%s' that includes:
		1. One-sentence overview of the package's main purpose
		2. 2-3 key functions/methods with simple examples
		3. Common use cases (1-2 sentences each)
		4. Gotchas or important notes beginners should know
		5. A minimal working example showing basic usage (10-15 lines max)
		Format everything in clear sections with headers. Use conversational language as if explaining to someone new to Go. 
		Focus on practical application rather than theory.`, prompt),
		Stream: true,
	}

	fmt.Println(request.Prompt)

	js, err := json.Marshal(&request)
	if err != nil {
		return err
	}

	client := http.Client{}
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:11434/api/generate", bytes.NewReader(js))
	if err != nil {
		return err
	}

	httpResp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	scanner := bufio.NewScanner(httpResp.Body)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var response Response
		if err := json.Unmarshal([]byte(line), &response); err != nil {
			return err
		}

		updateFn(response.Response, response.Done)

		if response.Done {
			break
		}
	}

	return scanner.Err()
}

func StreamAIExplanation(description string) tea.Cmd {
	fmt.Println(description)
	time.Sleep(2000 * time.Second)

	return func() tea.Msg {
		ch := make(chan tea.Msg)

		go func() {
			var content string
			content = "Loading AI explanation..."

			shared.Program.Send(AIStreamMsg{
				Content: content,
				Done: false,
			})

			content = ""
			
			err := Generate(description, func(fragment string, done bool) {
				content += fragment

				shared.Program.Send(AIStreamMsg{
					Content: content,
					Done: done,
				})
			})

			if err != nil {
				shared.Program.Send(AIStreamMsg{
					Error: err,
					Done: true,
				})
			}

			close(ch)
		}()

		return nil
	}
}
