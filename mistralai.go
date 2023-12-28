package mistralai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ResponseMessage struct {
	Content string
	Role    string
}

type ResponseChoices struct {
	FinishReason string
	Index        int
	Message      ResponseMessage
}

type ResponseData struct {
	ID      string
	Object  string
	Created float64
	Model   string
	Choices []ResponseChoices
	Usage   map[string]interface{}
}

func LoadConfig(fileName string) {
	err := godotenv.Load(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Request(msg string) {

	body := strings.NewReader(fmt.Sprintf(`
    {
    "model": "%s",
    "messages": [
     {
        "role": "%s",
        "content": "%s"
      }
    ],
    "safe_mode": true
  }`, os.Getenv("MODEL"), os.Getenv("ROLE"), msg))

	client := http.Client{}

	req, err := http.NewRequest(http.MethodPost, os.Getenv("ENDPOINT"), body)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	responseStruct := ResponseData{}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(responseBody, &responseStruct)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, result := range responseStruct.Choices {
		fmt.Println(result.Message.Content)
	}

	resp.Body.Close()
}
