package main

/*
Step 3: Needle in a haystack
------------------------------
We’re going to send you a dictionary with two values and keys.
The first value, needle, is a string. The second value, haystack, is an array of strings.
You’re going to tell the API where the needle is in the array.

Implementing this in Go was fun, but alot of work. I'm still learning the in's and out's of the language
*/

import (
	"fmt"
	"github.com/jochasinga/requests"
	"io/ioutil"
	"encoding/json"
)

func findNeedle(needle string, haystack []string) int {
	var temp int
	for key, val := range haystack {
		if needle == val {
			temp = key
		}
	}
	return temp
}

type Params struct {
	Token string `json:"token,omitempty"`
	Index int `json:"index,omitempty"`
}

type NeedleHayStack struct {
	Needle   string `json:"needle"`
	Haystack []string `json:"haystack"`
}

func main() {
	token := &Params{
		Token: "4f36c1b5089b07d3416239291874b956",
	}

	res, err := requests.PostJSON("http://challenge.code2040.org/api/haystack", token)
	if err != nil {
		panic(err)
	}

	// store as []byte
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	var test NeedleHayStack
	json.Unmarshal(page, &test)

	// fmt.Println(test)
	need := test.Needle
	hay := test.Haystack

	result := &Params {
		Token: "4f36c1b5089b07d3416239291874b956",
		Index: findNeedle(need, hay),
	}

	results, err := requests.PostJSON("http://challenge.code2040.org/api/haystack/validate", result)
	if err != nil {
		panic(err)
	}

	fmt.Println("index of needle: ", findNeedle(need, hay))
	fmt.Println(results.StatusCode) // post returns 400....

}
