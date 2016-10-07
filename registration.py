import requests

info = {'token': '4f36c1b5089b07d3416239291874b956', 'github': 'https://github.com/notdom/code2040'}
r = requests.post("http://challenge.code2040.org/api/register", json=info)

print(r.status_code)
