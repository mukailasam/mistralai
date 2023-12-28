## Mistralai

A simple to use MISTRAL AI client for interacting with Mistralai API

### Prerequisite

Make sure you have Golang installed and make sure to have a ".env" file in the project directory after cloning. the file content should look like this:

ENDPOINT=https://api.mistral.ai/v1/chat/completions \
API_KEY=Your Key \
MODEL=mistral-tiny \
ROLE=user

change the value of the API_KEY from "Your key" to your actual misralai API Key.
you can change the value of other element in the file but not must, only if you want to
 

### Install and RUN
`
- git  clone https://github.com/mukailasam/mistralai
- cd mistralai
- cd cmd
- go run main.go
`