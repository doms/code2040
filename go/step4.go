package main

/*
Step 4: Prefix
----------------
In this challenge, the API is going to give you another dictionary.
The first value, prefix, is a string. The second value, array, is an array of strings.
Your job is to return an array containing only the strings that do not start with that prefix.

Note: The strings in your array should be in the same order as in the original array.

Once I got the prefix and array, it was pretty easy just looping through the array and checking
each word to see if it was prefixed.
*/

import (
	"fmt"
	"github.com/jochasinga/requests"
	"io/ioutil"
	"encoding/json"

)


