import requests

'''
Step 4: Prefix
----------------
In this challenge, the API is going to give you another dictionary.
The first value, prefix, is a string. The second value, array, is an array of strings.
Your job is to return an array containing only the strings that do not start with that prefix.

Note: The strings in your array should be in the same order as in the original array.

Once I got the prefix and array, it was pretty easy just looping through the array and checking
each word to see if it was prefixed.
'''

def noPrefix(prefix, array):
	# list to hold non-prefixed words
	clean = []
	for word in array:
		if prefix not in word:
			clean.append(word)
	return clean

token = {'token':'4f36c1b5089b07d3416239291874b956'}
r = requests.post('http://challenge.code2040.org/api/prefix', json=token)

info = r.json()
prefix = info['prefix']
array = info['array']

# print("before prefix removal: ", array)
# print("prefix: ", prefix)
# print("after prefix removal: ", noPrefix(prefix, array))

result = {'token':'4f36c1b5089b07d3416239291874b956', 'array':noPrefix(prefix,array)}
s = requests.post('http://challenge.code2040.org/api/prefix/validate', json=result)
print(s.status_code)
