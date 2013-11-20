import json
import os
import re
import sys
import urllib
import urllib2

raw_books = [(x, open(os.path.join('books', x), 'r').read()) for x in os.listdir('./books')]
converted_books = []

header = re.compile('\*\*\*\s?START.*$', re.M)
footer = re.compile('\*\*\*\s?END.*$', re.M)
metadata = {
  'title': re.compile('Title:\s?(.*)$', re.M),
  'author': re.compile('Author:\s?(.*)$', re.M)
}

url = "http://localhost:8080/paragraphs"

for fname, book in raw_books:
  data = {}
  for n, r in metadata.iteritems():
    result = r.search(book)
    if result:
      data[n] = result.group(1).strip()

  content = footer.split(header.split(book)[1])[0]
  data['paragraphs'] = [x.strip() for x in content.split('\r\n\r\n') if len(x) > 0 and len(x.split('.')) > 1]
  data['id'] = int(fname[2:-4])
  print '%s:\t%s' % (data['id'], len(data['paragraphs']))

  req = urllib2.Request(url)
  req.add_header('Content-Type', 'application/json')

  try:
    resp = urllib2.urlopen(req, json.dumps(data))
  except Exception, e:
    print e.read()
