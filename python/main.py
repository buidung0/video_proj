import requests
import json

req = requests.post(
    'http://localhost:8080/videos',
    data=json.dumps({
        'title': 'hello world 2',
        'description': 'ipsum test 2',
        'url': 'http://localhost:8080/videos 111'
    }),
    headers={
        'Accept': 'application/json'
    }
)

print(req.text)
