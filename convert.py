import copy
import json
import re
import sys

import requests

raw_books = [(x, open(os.path.join('books', x), 'r').read()) for x in os.listdir('./books')]
converted_books = []

header = re.compile('\*\*\*\s?START.*$', re.M)
footer = re.compile('\*\*\*\s?END.*$', re.M)
metadata = {
  'title': re.compile('Title:\s?(.*)$', re.M),
  'author': re.compile('Author:\s?(.*)$', re.M)
}

for fname, book in raw_books:
  data = {}
  for n, r in metadata.iteritems():
    result = r.search(book)
    if result:
      data[n] = result.group(1).strip()

  content = footer.split(header.split(book)[1])[0]
  paragraphs = [x.strip() for x in content.strip().split('\r\n\r\n') if len(x) > 0]
  data['id'] = int(fname[2:-4])
  data['paragraphs'] = len(paragraphs)
  blob = copy.copy(data)
  blob['data'] = paragraphs

  print '%s:' % (data['id'], data['paragraphs'])

  try:
    r = requests.get("http://localhost:8080/upload")
    files = { 'file': json.dumps(blob) }
    r = requests.post(json.loads(r.text)['url'], files=files)

    data['key'] = json.loads(r.text)['key']
    r = requests.post("http://localhost:8080/books", json.dumps(data),
      headers={'Content-type': 'application/json'})
  except Exception, e:
    print e.read()
