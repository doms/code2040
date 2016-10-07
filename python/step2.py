import requests

'''
Step II: Reverse a string
---------------------------
The first one is straightforward. You’re going to reverse a string.
That is, if the API says “cupcake,” you’re going to send back “ekacpuc.”

I didn't realize reversing a string/list in Python was so easy!
'''

def reverseString(string):
	return string[::-1]

token = {'token':'4f36c1b5089b07d3416239291874b956'}
r = requests.post('http://challenge.code2040.org/api/reverse', json=token)
string = r.text

# print(reverseString(string))

result = {'token':'4f36c1b5089b07d3416239291874b956', 'string':reverseString(string)}
s = requests.post('http://challenge.code2040.org/api/reverse/validate', json=result)
print(s.status_code)
