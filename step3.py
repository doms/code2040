import requests

'''
Step 3: Needle in a haystack
------------------------------
We’re going to send you a dictionary with two values and keys.
The first value, needle, is a string. The second value, haystack, is an array of strings.
You’re going to tell the API where the needle is in the array.


This is probably the first time I've worked with dictionaries and their key/val pairs!
Once I found out how to store the val's in their respective variables, the rest was a piece of cake!
I also found out that Python has its own "indexOf()" function.
'''

def findNeedle(needle, haystack):
	for key in haystack:
		if key == needle:
			return haystack.index(needle)

token = {'token':'4f36c1b5089b07d3416239291874b956'}
r = requests.post('http://challenge.code2040.org/api/haystack', json=token)

info = r.json()
needle = info['needle']
haystack = info['haystack']

# print(findNeedle(needle, haystack))

result = {'token':'4f36c1b5089b07d3416239291874b956', 'needle':findNeedle(needle, haystack)}
s = requests.post('http://challenge.code2040.org/api/haystack/validate', json=result)
print(s.status_code)
