import requests
from datetime import date, timedelta
import dateutil.parser

'''
Stage 5: The dating game
----------------------------
The API will again give you a dictionary. The value for datestamp is a string, formatted as an ISO 8601 datestamp.
The value for interval is a number of seconds.

This challenge was very tricky for me. The only part of a datetime function I have used was to get the current day, so this was news to me about the other encodings.
I first wanted to find an easier way to read the iso8601 format and it turned out that there was a lib for it (yay!).
I figured I would first want to make both the datestamp and interval the same units, and since interval is presented in seconds,I considered it would be best to just convert the datestamp to seconds as well.
Then I went back to read the datetime documentation and low and behold, there was a way to convert it back to iso8601 format!
'''

def addInterval(datestamp, interval):
	# convert date to seconds and add interval
	date_in_seconds = datestamp + timedelta(seconds=interval)

	# convert back to iso8601 format
	new_date = date_in_seconds.isoformat('T')[0:19] + 'Z'
	return new_date


token = {'token':'4f36c1b5089b07d3416239291874b956'}
r = requests.post('http://challenge.code2040.org/api/dating', json=token)

info = r.json()

# datestamp = info['datestamp']
datestamp = dateutil.parser.parse(info['datestamp'])
interval = info['interval']

# print("original date: ", datestamp)
# print("interval: ", interval)
# print("new date: ", addInterval(datestamp, interval))

result = {'token':'4f36c1b5089b07d3416239291874b956', 'datestamp':addInterval(datestamp, interval)}
s = requests.post('http://challenge.code2040.org/api/dating/validate', json=result)
print(s.status_code)
