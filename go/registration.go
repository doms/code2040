package main

import (
	"fmt"
	"github.com/dghubble/sling"
)

// info = {"token": "4f36c1b5089b07d3416239291874b956", "github": "https://github.com/notdom/code2040"}

type Params struct {
	Token  string `json:"token,omitempty"`
	Github string `json:"github,omitempty"`
}

func main() {
	params := &Params{Token: "4f36c1b5089b07d3416239291874b956", Github: "https://github.com/notdom/code2040"}

	req, err := sling.New().Post("http://challenge.code2040.org/api/register").QueryStruct(params).Request()
	if err != nil {
		panic(err)
	}
	client.Do(req)
	fmt.Println(req)
}
