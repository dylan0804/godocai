package ai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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
	fmt.Println(prompt)

	request := Request{
        Model: "deepseek-r1:14b",
        Prompt: fmt.Sprintf(`Generate a practical explanation for the following Go type definition: %s`, prompt),
        Stream: true,
    }

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
