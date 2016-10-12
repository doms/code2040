package main

/*
Step II: Reverse a string
---------------------------
The first one is straightforward. You’re going to reverse a string.
That is, if the API says “cupcake,” you’re going to send back “ekacpuc.”
*/


import (
	"fmt"
	"github.com/jochasinga/requests"
	"io/ioutil"
)

// info = {"token": "4f36c1b5089b07d3416239291874b956", "github": "https://github.com/notdom/code2040"}

type Params struct {
	Token  string `json:"token,omitempty"`
	String string `json:"string,omitempty"`
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {

	token := &Params {
		Token: "4f36c1b5089b07d3416239291874b956",
	}

	res, err := requests.PostJSON("http://challenge.code2040.org/api/reverse",token)
	if err != nil {
		panic(err)
	}

	// store as []byte
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	// convert to string
	s := string(page)

	// fmt.Printf("%T \t%s \n", s, s)
	// fmt.Println("before reversed: ", s)
	// fmt.Println("after reversed: ", Reverse(s))

	result := &Params {
		Token: "4f36c1b5089b07d3416239291874b956",
		String: Reverse(s),
	}

	results, err := requests.PostJSON("http://challenge.code2040.org/api/reverse/validate", result)
	if err != nil {
		panic(err)
	}

	fmt.Println(results.StatusCode)
}
