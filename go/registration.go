package main

import (
	"fmt"
	"github.com/jochasinga/requests"
)

// info = {"token": "4f36c1b5089b07d3416239291874b956", "github": "https://github.com/notdom/code2040"}

type Params struct {
	Token  string `json:"token,omitempty"`
	Github string `json:"github,omitempty"`
}

func main() {

	data := &Params {
		Token: "4f36c1b5089b07d3416239291874b956",
		Github: "https://github.com/notdom/code2040",
	}

	res, err := requests.PostJSON("http://challenge.code2040.org/api/register",data)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.StatusCode) // should return 200
}
