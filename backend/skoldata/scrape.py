import requests

resp = requests.get('https://api.skolverket.se/skolenhetsregistret/v1/skolenhet')
skolenheter = resp.json()


f = open('skolenheter_detailed.jsonl', 'w+')

for skolenhet in skolenheter['Skolenheter']:
    kod = skolenhet['Skolenhetskod']
    print(kod)
    resp = requests.get('https://api.skolverket.se/skolenhetsregistret/v1/skolenhet/' + kod)

    if resp.status_code == 200:
        f.write(resp.text + '\n')
    else:
        print(kod, resp.text)