package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const OPENAI_URL = "http://localhost:11434/api/generate"

type RequestBody struct {
	Model  string
	Prompt string
	Stream bool
	Think  bool
}

type ResponseBody struct {
	Model      string
	CreatedAt  string
	Response   string
	Reason     string
	Done       bool
	DoneReason string
}

const PROMPT = `
You are an AI Kubernetes assistant.  
Constraints:  
- Be concise.  
- Output only valid YAML, kubectl commands, or short explanations when strictly needed.  
- The output format should be CLi friendly. Avoid Markdown and similar formattings.
- Prefer declarative manifests over imperative commands unless user requests otherwise.  
- Use stable API versions only.  
- Default namespace is "default" unless specified.  
- Assume recent Kubernetes LTS release.  
- Show complete code/config blocks, no truncation.  
- Do not add commentary or filler text.  
- Avoid hidden reasoning, output only the solution.
- Handle queries that are not related to kubernetes gracefully

User Question:
"%s"
`

func Run() {
	if len(os.Args) <= 1 { // first argument is name of program
		usage()
		os.Exit(1)
	}

	body := RequestBody{
		Model:  "huggingface.co/unsloth/Qwen3-30B-A3B-Instruct-2507-GGUF:latest",
		Prompt: fmt.Sprintf(PROMPT, os.Args[1]),
		Stream: false,
		Think:  false,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", OPENAI_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var resBodyJson ResponseBody
	err = json.Unmarshal(resBody, &resBodyJson)
	if err != nil {
		panic(err)
	}

	fmt.Println(resBodyJson.Response)
}

func usage() {
	fmt.Println("Usage: kubeai [QUESTION]")
}
